package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Task struct {
	ID          int
	Title       string
	Description string
	StartTime   time.Time
	EndTime     time.Time
	RepeatDays  string
	Cancel      bool
}

type ColorScheme struct {
	Base           rl.Color
	Surface        rl.Color
	Overlay        rl.Color
	Text           rl.Color
	Rose           rl.Color
	Subtle         rl.Color
	Muted          rl.Color
	Pine           rl.Color
	HighlightedLow rl.Color
	Gold           rl.Color
}
type GoBoardOpts struct {
	ColorScheme
	Font     rl.Font
	FontSize int32
}
type Goboard struct {
	DB           *sql.DB
	WindowHeight int
	WindowWidth  int

	GoBoardOpts
}

func NewGoboard(dbPath string, ww, wh int) (*Goboard, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := createTables(db); err != nil {
		return nil, err
	}

	opts, err := readConfig()
	if err != nil {
		log.Fatal("coudn't read config file")
	}
	return &Goboard{
		DB:           db,
		WindowWidth:  ww,
		WindowHeight: wh,
		GoBoardOpts:  opts,
	}, nil
}
func readConfig() (GoBoardOpts, error) {
	font := rl.LoadFont("./assets/fonts/MesloLGMNerdFont-Regular.ttf")
	return GoBoardOpts{
		ColorScheme: ColorScheme{
			Base:           rl.NewColor(25, 23, 36, 255),
			Surface:        rl.NewColor(31, 29, 46, 255),
			Overlay:        rl.NewColor(38, 35, 58, 255),
			Text:           rl.NewColor(224, 222, 244, 255),
			Rose:           rl.NewColor(235, 188, 186, 255),
			Subtle:         rl.NewColor(144, 140, 170, 255),
			Muted:          rl.NewColor(110, 106, 134, 255),
			Pine:           rl.NewColor(49, 116, 143, 255),
			HighlightedLow: rl.NewColor(33, 32, 46, 255),
			Gold:           rl.NewColor(246, 193, 119, 255),
		},
		Font:     font,
		FontSize: 20,
	}, nil
}

func createTables(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		start_time DATETIME ,
		end_time DATETIME ,
		repeat_days TEXT,
		cancel BOOLEAN DEFAULT FALSE
	);`
	_, err := db.Exec(query)
	return err
}

func (g *Goboard) Close() error {
	rl.UnloadFont(g.Font)
	return g.DB.Close()
}

func (g *Goboard) InsertTask(buffer []string) {
	query := `
	INSERT INTO tasks (title, description)
	VALUES (?, ?);`

	_, err := g.DB.Exec(query, buffer[0], buffer[1])
	if err != nil {
		log.Fatal("couldn't insert new task")
	}
}
func (g *Goboard) ReadTasks() {
	query := `
	SELECT id, title, description, start_time, end_time, repeat_days, cancel
	FROM tasks
	ORDER BY start_time;`

	rows, err := g.DB.Query(query)
	if err != nil {
		fmt.Errorf("failed to query tasks: %w", err)
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.StartTime,
			&task.EndTime,
			&task.RepeatDays,
			&task.Cancel,
		)
		if err != nil {
			fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, task)
	}

	// Check for errors after iteration
	if err := rows.Err(); err != nil {
		fmt.Errorf("error after iterating rows: %w", err)
	}

	for _, t := range tasks {
		fmt.Println(t.Title)
		fmt.Println(t.Description)
	}
}
