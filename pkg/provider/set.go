package provider

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/lib/pq"
	"github.com/tariqc80/oui-challenge/cmd/gqlgen/graph/model"
	"github.com/tariqc80/oui-challenge/internal/config"
)

type members []int64

func (a members) Len() int           { return len(a) }
func (a members) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a members) Less(i, j int) bool { return a[i] < a[j] }

// Set struct
type Set struct {
	db *sql.DB
}

// NewSet creates an instance of Set
func NewSet(c *config.Config) *Set {
	str := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.DatabaseUser, c.DatabasePassword, c.DatabaseHost, c.DatabasePort, c.DatabaseName)

	db, err := sql.Open("postgres", str)

	if err != nil {
		log.Print(err)
	}

	return &Set{
		db: db,
	}
}

// Close closes the database
func (s *Set) Close() {
	s.db.Close()
}

// Create inserts a new set
func (s *Set) Create(m members) (int64, error) {
	var (
		id           int64  // id of new set
		raw          []byte // byte array of members
		existingId   int64
		existingHash string
	)

	// sort the member array
	sort.Slice(m, m.Less)

	// convert members to byte array
	for _, num := range m {
		raw = strconv.AppendInt(raw, num, 10)
	}

	// get md5 hash of member data
	sum := md5.Sum(raw)

	// convert to a string so we can save it to the db
	hash := hex.EncodeToString(sum[:])

	q := "SELECT id, hash FROM sets WHERE hash = $1"
	err := s.db.QueryRow(q, hash).Scan(&existingId, &existingHash)
	if err != nil {
		log.Print(err)
	}

	if existingId != 0 {
		return 0, errors.New("Equivalent set already exists")
	}

	q = "INSERT INTO sets (members, hash) VALUES ($1, $2) RETURNING id"
	err = s.db.QueryRow(q, pq.Array([]int64(m)), hash).Scan(&id)
	if err != nil {
		log.Print(err)
	}

	return id, err
}

// Get fetches a set with the given id
func (s *Set) Get(id int64) (*model.Set, error) {
	var set model.Set
	var intersets string

	q := `
SELECT a.id, a.members, a.hash, jsonb_agg(b.id)
FROM sets AS a
INNER JOIN sets AS b ON (a.members && b.members)
WHERE a.id = $1
GROUP BY a.id
`

	row := s.db.QueryRow(q, id)

	err := row.Scan(&set.ID, pq.Array(&set.Members), &set.Hash, &intersets)

	if err != nil {
		return nil, err
	}

	// set.IntersectingSets =
	err = row.Err()

	if err != nil {
		log.Print(err)
	}

	return &set, err
}

// GetCollection queries the database for all the sets
func (s *Set) GetCollection() ([]*model.Set, error) {
	sets := []*model.Set{}

	q := `
SELECT a.id, a.members, a.hash, jsonb_agg(b.id)
FROM sets AS a
INNER JOIN sets AS b ON (a.members && b.members)
WHERE a.id <> b.id
GROUP BY a.id
`
	rows, err := s.db.Query(q)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var intersectionIds string

		currentSet := model.Set{}
		err := rows.Scan(&currentSet.ID, pq.Array(&currentSet.Members), &currentSet.Hash, &intersectionIds)

		if err != nil {
			return nil, err
		}

		var isets []*model.Set
		var intersects []interface{}

		err = json.Unmarshal([]byte(intersectionIds), &intersects)

		for _, id := range intersects {
			iset, err := s.Get(int64(id.(float64)))

			if err != nil {
				log.Print(err)
			}

			isets = append(isets, iset)
		}

		currentSet.IntersectingSets = isets

		sets = append(sets, &currentSet)
	}

	err = rows.Err()
	if err != nil {
		log.Print(err)
	}

	return sets, err
}
