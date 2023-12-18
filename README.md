# Management-System
<p>The House Rental System is a project made using gofr which helps to get the informations of the houses available for rent or are on sale.</p>

## Functions:
* **Get All Available Houses**: It is used to get the information of all the available houses.
* **Add New Home**: It can be used to add a new home in the database and stores all the information like address, price, description, availability for rent or it's for sale.
* **Get House Details By ID**: It is used to extract the information of a particular house using it's ID.
* **Update House Detail By ID**: It is used to update any details regarding the home like any change in price or in it's availability.
* **Delete a House By ID**: It is used when the owner doesn't wants its home to listed for availbility any more.

## :hammer_and_wrench: Installation:
    1. Cloning repository.
          git clone <repo link> or locally download zip folder.
          
    2. Set all enviorement variables in .env file.
          APP_NAME =  

    3. Add Your MongoDB URL in handler.go connectToDB() 

    4. Run following command to download and sync the project:
          go mod tidy
          go run main.go

    5. For Testing:
          go test
          
<h2>Sequence Diagram</h2>
<img src = "https://github.com/preritshrivastava2002/Management-System/assets/75294012/a9cc3e91-0b83-461e-84e4-09b1bfbca225">

<h2>Postman Collection</h2>
<p><a href = "https://github.com/preritshrivastava2002/Management-System/files/13698144/Postman.Test.Collection.json"> Postman Test Collection.json</p>
