# Management-System
<h2>The House Management System is an API where the following operations are implemented,</h2>
<ul>
  <li> CreateHome()  : if new home is available it can be inserted into the database with the address, description, city, state, postal code, available for rent, available for sale and it's price.</li>
  <li> AllHome() : This allows the users to see all the available houses and there complete information.</li>
  <li> GetHone() : This allows the users to fetch the information about a particular house from the database.</li>
  <li> UpdateHome()  : This allows to update any information regarding the house like change in price or it's availability for sale or rent.</li>
  <li> DeleteHome()  : If the house is to be removed from database then it will remove it.</li>
</ul>

## :hammer_and_wrench: Installation:
    1. Cloning repository.
          git clone <repo link> or locally download zip folder.
          
    2. Set all enviorement variables in .env file.
          APP_NAME =  

    3. Add Your MongoDB URL in handler.go coonectToDB() 

    4. Run command to update all the import with system
          go mod tidy
    
    5. Run the project using.
          go run main.go
          
<h2>Sequence Diagram</h2>

<h2>Postman Collection</h2>
<p><a href = "https://github.com/preritshrivastava2002/Management-System/files/13698144/Postman.Test.Collection.json"> Postman Test Collection.json</p>
