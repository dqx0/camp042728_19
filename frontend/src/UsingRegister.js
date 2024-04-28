import './App.css';
import React,{useState,useEffect} from "react"

export const UsingRegister=()=>{
    const [data,setData]=useState('');
    useEffect(()=>{
        fetch('http://localhost:8080/api/expense/1',{method:"GET"})
        .then((response)=>response.json)
        .then((getData)=>{
            setData(getData);
        })
    })
    return (
        <p>{data}</p>
    );
}
export default UsingRegister;