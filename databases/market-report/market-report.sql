BEGIN
    SELECT IFNULL(country, 'Total:') AS country, COUNT(competitor) AS competitors
    FROM foreignCompetitors
    GROUP BY country WITH ROLLUP;
END