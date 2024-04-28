/*
import './App.css';
import React,{useState,useEffect} from "react"

export const UsingRegister=()=>{
    const [data,setData]=useState('');
    useEffect(()=>{
        fetch('http://localhost:8080/api/expence/1',{method:"GET"})
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
*/
import './App.css';
import React, { useState, useEffect } from "react"

export const UsingRegister = () => {
    const [data, setData] = useState({});

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch('http://localhost:8080/api/expence/1', { method: "GET" });
                const jsonData = await response.json();
                setData(jsonData.data); // レスポンスの data オブジェクトをセット
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };

        fetchData();
    }, []);

    // 日付のフォーマット関数
    const formatDate = (dateString) => {
        const date = new Date(dateString);
        return date.toLocaleDateString(); // 日付をローカルの日付文字列に変換
    };

    return (
        <div>
            <p>ExpenseID: {data.ExpenseID}</p>
            <p>UserID: {data.UserID}</p>
            <p>Date: {formatDate(data.Date)}</p>
            <p>Title: {data.title}</p>
            <p>Money: {data.money}</p>
            <p>RemainingAmount: {data.RemainingAmount}</p>
        </div>
    );
};

export default UsingRegister;

