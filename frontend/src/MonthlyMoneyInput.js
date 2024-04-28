import './App.css';
import React, { useState } from "react"

export const MonthlyMoneyInput = () => {
    const [money, setMoney] = useState("");

    const handleChange = (e) => {
        setMoney(e.target.value);
    }

    const handleClick = () => {
        const currentDate = new Date();
        const year = currentDate.getFullYear();
        const month = currentDate.getMonth() + 1;

        console.log("handleClick発火");

        // ここで金額、年、月の順番でオブジェクトを作成
        const body = {
            money: parseInt(money, 10),
            year: year,
            month: month
        };

        fetch("http://localhost:8080/api/allowance", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(body)
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            console.log('Success:', data);
        })
        .catch(error => {
            console.error('Error:', error);
        });
    }

    return (
        <div>
            <p className='description'>１ヶ月に使用できる金額を入力してください</p>
            <p className='description'>金額</p>
            <div className='sub'>
                <input type='text' onChange={handleChange} value={money} />
            </div>
            <br></br>
            <div className='sub'>
                <button onClick={handleClick}>登録</button>
            </div>
        </div>
    );
}

export default MonthlyMoneyInput;