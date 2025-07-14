BEGIN
    BEGIN
	SELECT count(*) AS number_of_nulls
    FROM departments
        WHERE description REGEXP '^[[:space:]]*(NULL|nil|-)[[:space:]]*$' OR description IS NULL;
    END
END