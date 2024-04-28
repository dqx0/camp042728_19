import './App.css';
import React, { useState } from "react"

export const MonthlyMoneyInput = () => {
    const [money, setMoney] = useState("");

    const handleChange = (e) => {
        setMoney(e.target.value);
    }

    const handleClick = () => {
        console.log("handleClick発火");
        console.log("送信する金額: ", money);

        const body = {
            money: parseInt(money, 10)
        };

        // fetch APIを使用してサーバーにデータをPOSTする
        fetch(/*http://localhost:8080/api/allowance*/"http://localhost:3000", {
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

    // コンポーネントのUI部分
    return (
        <div>
            <p>金額</p>
            <input type='text' onChange={handleChange} value={money} />
            <button onClick={handleClick}>保存</button>
        </div>
    );
}

export default MonthlyMoneyInput;
