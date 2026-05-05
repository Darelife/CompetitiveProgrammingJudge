import sqlite3
from datetime import datetime

conn = sqlite3.connect("judge.db", check_same_thread=False)
cursor = conn.cursor()

cursor.execute("""
CREATE TABLE IF NOT EXISTS submissions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    code TEXT,
    status TEXT,
    output TEXT,
    created_at TEXT
)
""")
conn.commit()


def create_submission(code):
    cursor.execute(
        "INSERT INTO submissions (code, status, output, created_at) VALUES (?, ?, ?, ?)",
        (code, "PENDING", "", datetime.utcnow().isoformat())
    )
    conn.commit()
    return cursor.lastrowid


def update_submission(sub_id, status, output):
    cursor.execute(
        "UPDATE submissions SET status=?, output=? WHERE id=?",
        (status, output, sub_id)
    )
    conn.commit()


def get_submission(sub_id):
    cursor.execute("SELECT * FROM submissions WHERE id=?", (sub_id,))
    return cursor.fetchone()