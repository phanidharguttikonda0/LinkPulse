
CREATE TABLE Users (
                       id SERIAL PRIMARY KEY,
                       mail_id TEXT UNIQUE NOT NULL,
                       mobile VARCHAR(13) UNIQUE NOT NULL,
                       username VARCHAR(25) UNIQUE NOT NULL,
                       password VARCHAR(20) NOT NULL,
                       premium VARCHAR(1) NOT NULL DEFAULT '0', -- here 0 means free user , 1 means only for website urls, 2 for only file sharing, 3 for both
                       expiry TIMESTAMP,
                       totalspace_consumed FLOAT NOT NULL DEFAULT 0 -- IN MB
);

CREATE TABLE website_urls (
                              id SERIAL PRIMARY KEY,
                              user_id INT NOT NULL,
                              original_url TEXT NOT NULL,
                              shorten_url VARCHAR(25) NOT NULL,
                              clicks INT NOT NULL DEFAULT 0,
                              FOREIGN KEY (user_id) REFERENCES Users(id)
);

CREATE TABLE Files (
                       id SERIAL PRIMARY KEY,
                       user_id INT NOT NULL,
                       file_url TEXT NOT NULL,
                       space_taken FLOAT NOT NULL,
                       shared_url TEXT UNIQUE NOT NULL,
                       password VARCHAR(18),
                       mail_id BOOLEAN,
                       FOREIGN KEY (user_id) REFERENCES Users(id)
);

CREATE TABLE allowed (
                         id SERIAL PRIMARY KEY,
                         files_id INT NOT NULL,
                         allowed_mail_id TEXT NOT NULL,
                         FOREIGN KEY (files_id) REFERENCES Files(id)
);


CREATE TABLE payments (
                          id SERIAL PRIMARY KEY,
                          user_id INT NOT NULL REFERENCES Users(id) ON DELETE CASCADE,
                          transaction_id TEXT UNIQUE NOT NULL,
                          payment_gateway VARCHAR(30) NOT NULL,      -- e.g., Razorpay, Stripe
                          payment_status VARCHAR(20) NOT NULL,       -- e.g., SUCCESS, FAILED, PENDING
                          amount INT NOT NULL,                       -- in smallest unit (like paisa or cents)
                          currency VARCHAR(5) DEFAULT 'INR',         -- INR, USD, etc.
                          subscription_type VARCHAR(15) NOT NULL,    -- e.g., monthly, yearly, ultra
                          payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          expiry_date TIMESTAMP NOT NULL
);
