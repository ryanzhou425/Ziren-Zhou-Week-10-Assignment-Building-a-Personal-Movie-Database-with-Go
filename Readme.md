# Demo Overview
This project is built on the Go programming language and the SQLite database to create a localized personal movie database system. By importing the [IMDb dataset](https://arch.library.northwestern.edu/concern/datasets/3484zh40n?locale=en) provided by Northwestern University, it enables bulk storage and querying of movie information and genre data. The project uses a cgo-free driver to optimize performance and employs transactions and prepared statements to speed up data import. This system lays the foundation for future expansions such as personal movie collections, ratings, and recommendation features.

---
# Program Features
The core features of this project include:
- data import and query functionality 
- local database operations without relying on C language libraries 
- transaction processing to enhance import efficiency

---
# How to run
Use the following command to run the program:
- go run main.go

If everything works correctly, the local database movies.db will be automatically created during runtime, and movie information along with genre data will be imported in bulk.

The terminal will display the query results.

---
# SQL Example Query
The program executes an example SQL query to count the total number of movies in each genre, sorts them in descending order by quantity, and outputs the top ten genres.

**The results are as follows:**
- Genre: Short` ` ` ` ` `           Total Movies: 208263
- Genre: Drama` ` ` ` ` `           Total Movies: 185872
- Genre: Comedy` ` ` ` ` `          Total Movies: 144661
- Genre: Documentary` ` ` ` ` `     Total Movies: 108271
- Genre: Animation` ` ` ` ` `       Total Movies: 45175
- Genre: Action` ` ` ` ` `          Total Movies: 37260
- Genre: Romance` ` ` ` ` `         Total Movies: 34935
- Genre: Crime` ` ` ` ` `           Total Movies: 31559
- Genre: Family` ` ` ` ` `          Total Movies: 28774
- Genre: Thriller` ` ` ` ` `        Total Movies: 27017

---
# Design Concept for a Personal Collection Module
If I extend this database into a personal movie collection system, I would add a new table named personal_collection to store information about the movies collected by the user. Users would be able to systematically manage their watchlists, mark watched and to-watch movies, and review or filter content based on their own ratings and preferences. For example, the system could recommend similar movies based on the genres of highly rated films. Additionally, visualizations such as “Monthly Viewing Count” and “Most Frequently Collected Genres” could be generated to enhance user engagement and immersion.

---

# Future Directions and Potential Advantages of the Project
- A user account system can be introduced to support multi-user login and independent collection management.
- Features like “comments” and “tags” can be added, allowing users to give subjective feedback on films and assign personalized labels.
- Using machine learning algorithms, the system can generate personalized recommendation lists based on user ratings and viewing history.

Compared to IMDb, a personal movie database emphasizes privacy, customization, and local control. Users can tailor their viewing experience according to their own preferences without being restricted by the limitations of the IMDb platform. This flexibility is especially valuable for film enthusiasts, critics, or users seeking to build a long-term viewing archive.

---

# Use of AI Assistants
- Searched for important considerations when writing SQL statements in Go
- Searched for why the program remains unresponsive for a long time after running go run main.go
- Searched for how to speed up bulk data insertion in Go
- Searched for what a transaction is in and why it can significantly improve performance


---