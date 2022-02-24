CREATE DATABASE IF NOT EXISTS jobs;

use jobs;

CREATE TABLE IF NOT EXISTS companies (
  company_id INT NOT NULL AUTO_INCREMENT,
  company VARCHAR(255) NOT NULL,
  PRIMARY KEY (company_id)
);

CREATE TABLE IF NOT EXISTS users (
  user_id INT NOT NULL AUTO_INCREMENT,
  username VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  role VARCHAR(255) NOT NULL,
  company_id INT NOT NULL,
  PRIMARY KEY (user_id),
  FOREIGN KEY (company_id) REFERENCES companies(company_id)
);

CREATE TABLE IF NOT EXISTS profiles (
  profile_id INT NOT NULL AUTO_INCREMENT,
  profile VARCHAR(255) NOT NULL,
  PRIMARY KEY (profile_id)
);

CREATE TABLE IF NOT EXISTS job_states (
  state_id INT NOT NULL AUTO_INCREMENT,
  state VARCHAR(255) NOT NULL,
  PRIMARY KEY (state_id)
);

CREATE TABLE IF NOT EXISTS job_cities (
  city_id INT NOT NULL AUTO_INCREMENT,
  city VARCHAR(255) NOT NULL,
  state_id INT NOT NULL,
  PRIMARY KEY (city_id),
  FOREIGN KEY (state_id) REFERENCES job_states(state_id)
);

CREATE TABLE IF NOT EXISTS jobs (
  job_id INT NOT NULL AUTO_INCREMENT,
  title VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  status VARCHAR(255) NOT NULL,
  user_id INT NOT NULL,
  city_id INT NOT NULL,
  state_id INT NOT NULL,

  PRIMARY KEY (job_id),
  FOREIGN KEY (user_id) REFERENCES users(user_id),
  FOREIGN KEY (city_id) REFERENCES job_cities(city_id),
  FOREIGN KEY (state_id) REFERENCES job_states(state_id)
);

CREATE TABLE IF NOT EXISTS jobs_profiles (
  profile_id INT NOT NULL AUTO_INCREMENT,
  job_id INT NOT NULL,
  FOREIGN KEY (job_id) REFERENCES jobs(job_id),
  FOREIGN KEY (profile_id) REFERENCES profiles(profile_id)
);

-- SEED DATA

INSERT INTO companies (company) VALUES ('Company 1');
INSERT INTO companies (company) VALUES ('Company 2');
INSERT INTO companies (company) VALUES ('Company 3');

INSERT INTO users (username, password, role, company_id) VALUES ('admin', 'admin', 'admin', 1);
INSERT INTO users (username, password, role, company_id) VALUES ('user', 'user', 'recruiter', 2);
INSERT INTO users (username, password, role, company_id) VALUES ('user2', 'user', 'recruiter', 3);

INSERT INTO profiles (profile) VALUES ('CLT');
INSERT INTO profiles (profile) VALUES ('PJ');
INSERT INTO profiles (profile) VALUES ('Estágio');

INSERT INTO job_states (state) VALUES ("SP");
INSERT INTO job_states (state) VALUES ("SC");
INSERT INTO job_states (state) VALUES ("RS");

INSERT INTO job_cities (city, state_id) VALUES ("Santo André", 1);
INSERT INTO job_cities (city, state_id) VALUES ("Floripa", 2);
INSERT INTO job_cities (city, state_id) VALUES ("Não me toque", 3);

INSERT INTO jobs (title,description,status,user_id,city_id,state_id) 
VALUES ("Desenvolvedor PHP", "Vaga para desenvolvedor PHP", "A", 1, 1, 1);

INSERT INTO jobs (title,description,status,user_id,city_id,state_id) 
VALUES ("Desenvolvedor Web3", "Vaga para desenvolvedor Web3", "I", 2, 2, 2);

INSERT INTO jobs (title,description,status,user_id,city_id,state_id) 
VALUES ("Desenvolvedor Golang", "Vaga para desenvolvedor Golang", "V", 3, 3, 3);

INSERT INTO jobs (title,description,status,user_id,city_id,state_id) 
VALUES ("Desenvolvedor JS", "Vaga para desenvolvedor JS", "V", 1, 2, 2);

INSERT INTO jobs (title,description,status,user_id,city_id,state_id) 
VALUES ("Desenvolvedor Web3", "Vaga para desenvolvedor Web3", "A", 2, 3, 3);

INSERT INTO jobs (title,description,status,user_id,city_id,state_id) 
VALUES ("Desenvolvedor Python", "Vaga para desenvolvedor Python", "I", 3, 1, 1);

INSERT INTO jobs (title,description,status,user_id,city_id,state_id) 
VALUES ("Desenvolvedor Java", "Vaga para desenvolvedor Java", "A", 1, 1, 1);

INSERT INTO jobs_profiles (job_id, profile_id) VALUES (1, 1);
INSERT INTO jobs_profiles (job_id, profile_id) VALUES (2, 2);
INSERT INTO jobs_profiles (job_id, profile_id) VALUES (3, 3);
INSERT INTO jobs_profiles (job_id, profile_id) VALUES (4, 1);
INSERT INTO jobs_profiles (job_id, profile_id) VALUES (5, 2);
INSERT INTO jobs_profiles (job_id, profile_id) VALUES (6, 3);