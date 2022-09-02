package db

import (
	"context"
	"database/sql"
	"github.com/ahmedkhaeld/bookings-api/internals/models"
	"github.com/ahmedkhaeld/bookings-api/internals/repository"
	"log"
	"time"
)

// Postgres implements the Bookings interface by sql methods
type Postgres struct {
	DB *sql.DB
}

// PostgresRepo construct a postgres connection
func PostgresRepo(conn *sql.DB) repository.Bookings {
	return &Postgres{
		DB: conn,
	}
}

//FetchAllRooms returns a collection of stored rooms
func (p *Postgres) FetchAllRooms() ([]*models.Room, error) {
	//use the context because the user might lose connection for any reason
	// this track the request within a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var rooms []*models.Room

	query := `select id, room_name, price, rate, created_at, updated_at from rooms order by room_name`

	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Room
		err := rows.Scan(
			&r.ID,
			&r.RoomName,
			&r.Price,
			&r.Rate,
			&r.CreatedAt,
			&r.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, &r)
	}
	if err = rows.Err(); err != nil {
		return rooms, err
	}

	return rooms, nil
}

//FetchRoom returns a room by a given id
func (p *Postgres) FetchRoom(id int) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// query the wanted movie by its id
	query := `select id, room_name, price, rate, created_at, updated_at from rooms where id = $1`

	var room models.Room
	row := p.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&room.ID,
		&room.RoomName,
		&room.Price,
		&room.Rate,
		&room.CreatedAt,
		&room.UpdatedAt,
	)
	if err != nil {
		return room, err
	}
	return room, nil
}

//UpdateRoom modify a room data record by id
func (p *Postgres) UpdateRoom(r models.Room) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update rooms set room_name = $1, price=$2, rate=$3, updated_at=$4 where id = $5`

	_, err := p.DB.ExecContext(ctx, query,
		r.RoomName,
		r.Price,
		r.Rate,
		time.Now(),
		r.ID,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

//InsertRoom add new record in rooms table
func (p *Postgres) InsertRoom(r models.Room) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into rooms ( id,room_name,price, rate, created_at, updated_at) values ($1, $2,$3,$4,$5,$6)`

	_, err := p.DB.ExecContext(ctx, query,
		r.ID,
		r.RoomName,
		r.Price,
		r.Rate,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//DeleteRoom removes a room record by id
func (p *Postgres) DeleteRoom(r models.Room) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `delete from rooms where id = $1`

	_, err := p.DB.ExecContext(ctx, query, r.ID)
	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var rooms []models.Room

	query := `
		select 
			r.id, r.room_name, price, rate, updated_at
		from
			rooms r
		where r.id not in 
			(select 
				rr.room_id 
			from 
				room_restrictions rr 
			where 
			$1 < rr.end_date and $2 > rr.start_date)`

	rows, err := m.DB.QueryContext(ctx, query, start, end)
	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		var room models.Room

		err := rows.Scan(
			&room.ID,
			&room.RoomName,
			&room.Price,
			&room.Rate,
			&room.UpdatedAt,
		)
		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)
	}
	if err = rows.Err(); err != nil {
		return rooms, err
	}
	return rooms, nil
}

func (m *Postgres) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var numRows int
	query := `
			select 
				count(id) 
			from 
				room_restrictions
			where 
			      room_id = $1
				and ($2 < end_date and $3 > start_date)`

	row := m.DB.QueryRowContext(ctx, query, roomID, start, end)
	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}

	if numRows == 0 {
		return true, nil
	}
	return false, nil
}
