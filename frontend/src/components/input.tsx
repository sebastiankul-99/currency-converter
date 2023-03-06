import react, {useEffect, useState} from 'react';
import Form from 'react-bootstrap/Form';
import InputGroup from 'react-bootstrap/InputGroup';
import {Currency} from "../types/currency";
import {BackedHost} from "../types/enviroments";


const getRate = async (code:string, setRate: (t:number) => void)=>{

    if(code==="PLN"){
        setRate(1);
        return;
    }
    await fetch(`${BackedHost.HOST}:${BackedHost.PORT}/rate/${code}`)
        .then((response) => response.json())
        .then((content) =>{
            setRate(Number(content.rate));
        });

}

export const CustomInput = (props: {currencies: Array<Currency>}) => {

    const [firstCurrency, setFirstCurrency] = useState<string>("PLN");
    const [secondCurrency, setSecondCurrency] = useState<string>("PLN");
    const [current,setCurrent] = useState<number|null>(null);
    const [firstValue, setFirstValue] = useState<string>("");
    const [secondValue, setSecondValue] = useState<string>("");
    const [firstRate, setFirstRate] = useState<number>(1);
    const [secondRate, setSecondRate] = useState<number>(1);

    useEffect(() => {
        const timer = setTimeout(() => {
            if(firstValue!=="" && (current===1|| current ===null)){
                let value = Math.round(Number(firstValue) * firstRate * (1/secondRate)*10000)/10000;
                //rounding to 4 decimal places
                setSecondValue(value.toString());
            }
        }, 350);

        return () => clearTimeout(timer);
    }, [firstValue])

    useEffect(() => {
        const secondTimer = setTimeout(() => {
            if(secondValue!=="" && (current === 2 || current ===null)){
                let value =Math.round( Number(secondValue) * secondRate * (1/firstRate)*10000)/10000;
                setFirstValue(value.toString());
            }
        }, 350);

        return () => clearTimeout(secondTimer);
    }, [secondValue])

    useEffect(()=>{
        getRate(
            firstCurrency,
            setFirstRate
        ).catch(error => console.log(error));

    },[firstCurrency])

    useEffect(()=>{
        getRate(
            secondCurrency,
            setSecondRate
        ).catch(error => console.log(error));
    },[secondCurrency])

    useEffect(()=> {
        if(current===1){
            let value = Math.round(Number(firstValue) * firstRate * (1/secondRate)*10000)/10000;
            setSecondValue(value.toString());
        } else {
            let value = Math.round(Number(secondValue) * secondRate * (1/firstRate)*10000)/10000;
            setFirstValue(value.toString());
        }

    },[firstRate, secondRate])


    return(
    <div>
        <InputGroup className="mb-3">
            <Form.Control value={firstValue.toString()}
                          onChange={(e) => {
                              let tmp = Number(e.target.value);
                              if(!isNaN(tmp) && tmp>=0) {
                                  setFirstValue(e.target.value);
                              }
                          }}
                          onClick={(e) => setCurrent(1)}
            />
            <Form.Select defaultValue={"USD A"} onChange={(e) => setFirstCurrency(e.target.value)}>
                {
                    props.currencies.map(
                        (currency)=>(
                            <option value = {currency.code} key = {currency.code}>
                                {currency.name}
                            </option>
                        )
                    )
                }
            </Form.Select>
        </InputGroup>
        <InputGroup className="mb-3">
            <Form.Control value={secondValue.toString()}
                          onChange={(e) => {
                              let tmp = Number(e.target.value);
                              if(!isNaN(tmp) && tmp>=0) {
                                  setSecondValue(e.target.value);
                              }
                          }}
                          onClick={(e) => setCurrent(2)}
            />
            <Form.Select onChange={(e) => setSecondCurrency(e.target.value)}>
                {
                    props.currencies.map(
                        (currency)=>(
                            <option value = {currency.code} key = {currency.code}>
                                {currency.name}
                            </option>
                        )
                    )
                }
            </Form.Select>
        </InputGroup>
    </div>
    )
}






