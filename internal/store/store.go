package store
import ("database/sql";"fmt";"os";"path/filepath";"time";_ "modernc.org/sqlite")
type DB struct{db *sql.DB}
type Channel struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Type string `json:"type"`
	MemberCount int `json:"member_count"`
	MessageCount int `json:"message_count"`
	LastMessageAt string `json:"last_message_at"`
	Status string `json:"status"`
	CreatedAt string `json:"created_at"`
}
func Open(d string)(*DB,error){if err:=os.MkdirAll(d,0755);err!=nil{return nil,err};db,err:=sql.Open("sqlite",filepath.Join(d,"saloon.db")+"?_journal_mode=WAL&_busy_timeout=5000");if err!=nil{return nil,err}
db.Exec(`CREATE TABLE IF NOT EXISTS channels(id TEXT PRIMARY KEY,name TEXT NOT NULL,description TEXT DEFAULT '',type TEXT DEFAULT 'public',member_count INTEGER DEFAULT 0,message_count INTEGER DEFAULT 0,last_message_at TEXT DEFAULT '',status TEXT DEFAULT 'active',created_at TEXT DEFAULT(datetime('now')))`)
return &DB{db:db},nil}
func(d *DB)Close()error{return d.db.Close()}
func genID()string{return fmt.Sprintf("%d",time.Now().UnixNano())}
func now()string{return time.Now().UTC().Format(time.RFC3339)}
func(d *DB)Create(e *Channel)error{e.ID=genID();e.CreatedAt=now();_,err:=d.db.Exec(`INSERT INTO channels(id,name,description,type,member_count,message_count,last_message_at,status,created_at)VALUES(?,?,?,?,?,?,?,?,?)`,e.ID,e.Name,e.Description,e.Type,e.MemberCount,e.MessageCount,e.LastMessageAt,e.Status,e.CreatedAt);return err}
func(d *DB)Get(id string)*Channel{var e Channel;if d.db.QueryRow(`SELECT id,name,description,type,member_count,message_count,last_message_at,status,created_at FROM channels WHERE id=?`,id).Scan(&e.ID,&e.Name,&e.Description,&e.Type,&e.MemberCount,&e.MessageCount,&e.LastMessageAt,&e.Status,&e.CreatedAt)!=nil{return nil};return &e}
func(d *DB)List()[]Channel{rows,_:=d.db.Query(`SELECT id,name,description,type,member_count,message_count,last_message_at,status,created_at FROM channels ORDER BY created_at DESC`);if rows==nil{return nil};defer rows.Close();var o []Channel;for rows.Next(){var e Channel;rows.Scan(&e.ID,&e.Name,&e.Description,&e.Type,&e.MemberCount,&e.MessageCount,&e.LastMessageAt,&e.Status,&e.CreatedAt);o=append(o,e)};return o}
func(d *DB)Update(e *Channel)error{_,err:=d.db.Exec(`UPDATE channels SET name=?,description=?,type=?,member_count=?,message_count=?,last_message_at=?,status=? WHERE id=?`,e.Name,e.Description,e.Type,e.MemberCount,e.MessageCount,e.LastMessageAt,e.Status,e.ID);return err}
func(d *DB)Delete(id string)error{_,err:=d.db.Exec(`DELETE FROM channels WHERE id=?`,id);return err}
func(d *DB)Count()int{var n int;d.db.QueryRow(`SELECT COUNT(*) FROM channels`).Scan(&n);return n}

func(d *DB)Search(q string, filters map[string]string)[]Channel{
    where:="1=1"
    args:=[]any{}
    if q!=""{
        where+=" AND (name LIKE ? OR description LIKE ?)"
        args=append(args,"%"+q+"%");args=append(args,"%"+q+"%");
    }
    if v,ok:=filters["type"];ok&&v!=""{where+=" AND type=?";args=append(args,v)}
    if v,ok:=filters["status"];ok&&v!=""{where+=" AND status=?";args=append(args,v)}
    rows,_:=d.db.Query(`SELECT id,name,description,type,member_count,message_count,last_message_at,status,created_at FROM channels WHERE `+where+` ORDER BY created_at DESC`,args...)
    if rows==nil{return nil};defer rows.Close()
    var o []Channel;for rows.Next(){var e Channel;rows.Scan(&e.ID,&e.Name,&e.Description,&e.Type,&e.MemberCount,&e.MessageCount,&e.LastMessageAt,&e.Status,&e.CreatedAt);o=append(o,e)};return o
}

func(d *DB)Stats()map[string]any{
    m:=map[string]any{"total":d.Count()}
    rows,_:=d.db.Query(`SELECT status,COUNT(*) FROM channels GROUP BY status`)
    if rows!=nil{defer rows.Close();by:=map[string]int{};for rows.Next(){var s string;var c int;rows.Scan(&s,&c);by[s]=c};m["by_status"]=by}
    return m
}
