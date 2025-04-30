BEGIN
    SELECT name FROM leaderboard;
    ORDER BY score DES LIMIT 3,5;
END