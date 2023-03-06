import React, {useEffect, useState} from 'react';
import {CustomInput} from './components/input';
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import {Currency} from "./types/currency";
import {BackedHost} from "./types/enviroments";
function App() {

    const [currencyTable, setCurrencyTable] = useState<Array<Currency>>([])
    useEffect(() => {
        let tempTable:Array<Currency> = [{
            code: "PLN",
            name: "złoty"
        }];

        (
            async function (){
                await fetch(`${BackedHost.HOST}:${BackedHost.PORT}/currencies`, {

                }).then(response => response.json()
                ).then((content:Array<Currency>) =>{
                    tempTable = tempTable.concat(content)
                    setCurrencyTable(tempTable)
                }).catch(error => console.log(error));
            }
        )();

    }, []);

  return (

    <div className="App">
      <header className="App-header">
        <CustomInput currencies={currencyTable} />
      </header>
    </div>
  );
}

export default App;
