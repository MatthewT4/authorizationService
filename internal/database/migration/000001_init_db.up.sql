CREATE TABLE "users" (
                                "user_id" serial NOT NULL,
                                "email" varchar(320) NOT NULL UNIQUE,
                                "password" varchar(255) NOT NULL,
                                "name" varchar(200) NOT NULL,
                                "surname" varchar(200) NOT NULL,
                                "phone" varchar(50) NOT NULL,
                                CONSTRAINT "users_pk" PRIMARY KEY ("user_id")
) WITH (
      OIDS=FALSE
      );



CREATE TABLE "session" (
                                  "user_id" integer NOT NULL,
                                  "service_id" integer NOT NULL,
                                  "create_date" DATE NOT NULL,
                                  "valid_until_date" DATE NOT NULL,
                                  "session_token" varchar(50) NOT NULL,
                                  "session_id" integer NOT NULL
) WITH (
      OIDS=FALSE
      );



CREATE TABLE "service" (
                                  "service_id" serial NOT NULL,
                                  "user_id" integer NOT NULL,
                                  "name" varchar(150) NOT NULL,
                                  "domains" varchar(300) NOT NULL,
                                  "description" TEXT NOT NULL,
                                  "private_key" varchar(100) NOT NULL,
                                  CONSTRAINT "service_pk" PRIMARY KEY ("service_id")
) WITH (
      OIDS=FALSE
      );




ALTER TABLE "session" ADD CONSTRAINT "session_fk0" FOREIGN KEY ("user_id") REFERENCES "users"("user_id");
ALTER TABLE "session" ADD CONSTRAINT "session_fk1" FOREIGN KEY ("service_id") REFERENCES "service"("service_id");

ALTER TABLE "service" ADD CONSTRAINT "service_fk0" FOREIGN KEY ("user_id") REFERENCES "users"("user_id");