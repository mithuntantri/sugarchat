CREATE TABLE IF NOT EXISTS "users" (
  "mobile_number" varchar(10) NOT NULL,
  "email_id" varchar(127) NOT NULL,
  "firstname" varchar(30) NOT NULL,
  "lastname" varchar(30),
  "dob" TIMESTAMP WITH TIME ZONE NOT NULL,
  "gender_code" varchar(1) NOT NULL,
  CONSTRAINT profile_pk PRIMARY KEY ("mobile_number")
) WITH (
  OIDS=FALSE
)
