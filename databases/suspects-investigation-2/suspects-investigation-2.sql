BEGIN
    SELECT id, name, surname FROM Suspect;
    WHERE heigh <= 170 AND 
    SUBSTRING(name, 1, 1) <> 'B'
    AND surname NOT LIKE 'Gre_n'
    ORDER BY id;
END