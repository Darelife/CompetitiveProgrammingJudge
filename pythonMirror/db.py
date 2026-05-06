import sqlite3
from datetime import datetime

conn = sqlite3.connect("judge.db", check_same_thread=False)
cursor = conn.cursor()

cursor.execute("""
CREATE TABLE IF NOT EXISTS submissions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    code TEXT,
    language TEXT,
    question_id TEXT,
    status TEXT,
    output TEXT,
    created_at TEXT
)
""")
conn.commit()


def create_submission(code, language, question_id):
    cursor.execute(
        """
        INSERT INTO submissions
        (code, language, question_id, status, output, created_at)
        VALUES (?, ?, ?, ?, ?, ?)
        """,
        (
            code,
            language,
            question_id,
            "PENDING",
            "",
            datetime.utcnow().isoformat()
        )
    )

    conn.commit()
    return cursor.lastrowid


def update_submission(sub_id, status, output):
    # Check for sql injection
    cursor.execute(
        "UPDATE submissions SET status=?, output=? WHERE id=?",
        (status, output, sub_id)
    )
    conn.commit()


def get_submission(sub_id):
    # Check for SQL injection
    cursor.execute("SELECT * FROM submissions WHERE id=?", (sub_id,))
    return cursor.fetchone()
