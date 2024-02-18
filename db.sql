-- Create Users Table
CREATE TABLE users (
    userid SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);

-- Create Conversations Table
CREATE TABLE conversations (
    conversationid SERIAL PRIMARY KEY,
    user1id INT NOT NULL,
    user2id INT NOT NULL,
    createdat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user1id) REFERENCES users(userid),
    FOREIGN KEY (user2id) REFERENCES users(userid)
);

-- Create Messages Table
CREATE TABLE messages (
    messageid SERIAL PRIMARY KEY,
    conversationid INT NOT NULL,
    senderuserid INT NOT NULL,
    receiveruserid INT NOT NULL,
    content TEXT,
    sentat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (conversationid) REFERENCES conversations(conversationid),
    FOREIGN KEY (senderuserid) REFERENCES users(userid),
    FOREIGN KEY (receiveruserid) REFERENCES users(userid)
);
