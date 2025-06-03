BEGIN
    SELECT COUNT(*), AVG(population) as average, SUM(population) as total
    FROM countries; 
END