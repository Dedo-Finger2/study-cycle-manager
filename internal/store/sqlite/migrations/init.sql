CREATE TABLE IF NOT EXISTS study_cycles (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT NOT NULL UNIQUE,
  completed_times INTEGER NOT NULL,
  selected BOOLEAN NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS study_cycle_subjects (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  study_cycle_id INTEGER NOT NULL,
  name TEXT NOT NULL,
  max_study_hours INTEGER NOT NULL,
  user_studied_hours INTEGER NOT NULL,
  completed_times INTEGER NOT NULL,
  added_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,

  FOREIGN KEY (study_cycle_id) REFERENCES study_cycles(id)
  UNIQUE (study_cycle_id, name)
);
