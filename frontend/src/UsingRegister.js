import './App.css';
import React,{useState} from "react"

export const UsingRegister=()=>{
    const[use,setInputUse]=useState("");
    const handleChange=(e)=>{
        setInputUse(e.target.value);
        console.log(use+":error")
    };
    const handleClick=(e)=>{
        console.log("aaaa");
        const body={
            use:use
        };
        fetch("http://localhost:/3000/use",{
            method:'POST',
            body:JSON.stringify(body)
        })
        .then(response=>{
            if(!response.ok){
                throw new Error('response error');
            }
            console.log(body);
            return response.json();
        })
        .then(data=>{
            console.log(data);
        })
    }
    return (
        <div>
            <p>金額</p>
            <input onClick={handleClick} onChange={handleChange} />
            <button onClick={handleClick}>送信</button>
        </div>
    );
}
export default UsingRegister;