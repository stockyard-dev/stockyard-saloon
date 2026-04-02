package store
import ("database/sql";"fmt";"os";"path/filepath";"time";_ "modernc.org/sqlite")
type DB struct{db *sql.DB}
type Message struct{
	ID string `json:"id"`
	Channel string `json:"channel"`
	Author string `json:"author"`
	Body string `json:"body"`
	ReplyTo string `json:"reply_to"`
	CreatedAt string `json:"created_at"`
}
func Open(d string)(*DB,error){if err:=os.MkdirAll(d,0755);err!=nil{return nil,err};db,err:=sql.Open("sqlite",filepath.Join(d,"saloon.db")+"?_journal_mode=WAL&_busy_timeout=5000");if err!=nil{return nil,err}
db.Exec(`CREATE TABLE IF NOT EXISTS messages(id TEXT PRIMARY KEY,channel TEXT DEFAULT 'general',author TEXT NOT NULL,body TEXT DEFAULT '',reply_to TEXT DEFAULT '',created_at TEXT DEFAULT(datetime('now')))`)
return &DB{db:db},nil}
func(d *DB)Close()error{return d.db.Close()}
func genID()string{return fmt.Sprintf("%d",time.Now().UnixNano())}
func now()string{return time.Now().UTC().Format(time.RFC3339)}
func(d *DB)Create(e *Message)error{e.ID=genID();e.CreatedAt=now();_,err:=d.db.Exec(`INSERT INTO messages(id,channel,author,body,reply_to,created_at)VALUES(?,?,?,?,?,?)`,e.ID,e.Channel,e.Author,e.Body,e.ReplyTo,e.CreatedAt);return err}
func(d *DB)Get(id string)*Message{var e Message;if d.db.QueryRow(`SELECT id,channel,author,body,reply_to,created_at FROM messages WHERE id=?`,id).Scan(&e.ID,&e.Channel,&e.Author,&e.Body,&e.ReplyTo,&e.CreatedAt)!=nil{return nil};return &e}
func(d *DB)List()[]Message{rows,_:=d.db.Query(`SELECT id,channel,author,body,reply_to,created_at FROM messages ORDER BY created_at DESC`);if rows==nil{return nil};defer rows.Close();var o []Message;for rows.Next(){var e Message;rows.Scan(&e.ID,&e.Channel,&e.Author,&e.Body,&e.ReplyTo,&e.CreatedAt);o=append(o,e)};return o}
func(d *DB)Delete(id string)error{_,err:=d.db.Exec(`DELETE FROM messages WHERE id=?`,id);return err}
func(d *DB)Count()int{var n int;d.db.QueryRow(`SELECT COUNT(*) FROM messages`).Scan(&n);return n}
