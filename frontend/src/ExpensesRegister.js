import './App.css';
import React, { useState } from "react"


export const ExpensesRegister = () => {
    // テキストボックスに入力された値をtext変数に保存
    const [title, setTitle] = useState();
    const [money, setMoney] = useState();
    let count = 0;

    // テキストボックスの中身が変更された時にonChangeが実行される
    // onChange={handleChange}のようにすると、
    // onChangeしたときにhandleChangeが実行される
    const handleChange = (e) => {
        setTitle(e.target.value);
        console.log(title);
    }

    const handleChange2 = (e) => {
        setMoney(e.target.value);
        console.log(money);
    }

    const handleClick = (e) => {
        console.log("handleClick発火");
        console.log(title);
        console.log(money);

        const body = {
            title: title,
            money: parseInt(money)
        }
        if (count < 1) { // 新しい支出を記録します.
            // ローカル環境でのテスト用に3000番ポートのurlにしています。
            fetch(/*"http://localhost:8080/api/expenses"*/"http://localhost:3000", {
                method: 'POST',
                body: JSON.stringify(body)
            })
            .then(
                console.log(body)
            )
            .then(
                count++
            )
            .then(
                console.log("count: " + count)
            );
        } else { // 既存の支出を記録する際に使用します.
            // ローカル環境でのテスト用に3000番ポートのurlにしています。
            fetch(/*"http://localhost:8080/api/expenses"*/"http://localhost:3000", {
                method: 'PUT',
                body: JSON.stringify(body)
            })
            .then(
                console.log(body)
            )
            .then(
                count++
            );
        }
    
    }

    return (
        <div>
            {/* テキストボックス */}
            <p>商品名</p>
            <input type='text' onChange={handleChange} />
            <p>金額</p>
            <input type='text' onChange={handleChange2} />

            {/* 保存ボタン */}
            <button onClick={handleClick}>保存</button>
        </div>
    );
}

export default ExpensesRegister;
