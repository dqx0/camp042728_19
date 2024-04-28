import './App.css';
import React, { useState } from "react"
import Header from './Header';


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
        console.log(title);
        console.log(money);

        const body = {
            title: title,
            money: parseInt(money)
        }
        // 新しい支出を記録します.
        // ローカル環境でのテスト用に3000番ポートのurlにしています。
        fetch("http://localhost:8080/api/expence"/*"http://localhost:3000"*/, {
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
    
    }

    return (
        <div>
            {/* <Header /> */}
            {/* テキストボックス */}
            <p className='description'>商品名</p>
            <div className='sub'>

                <input type='text' onChange={handleChange} className='textbox'/>
            </div>
            <p className='description'>金額</p>
            <div className='sub'>
            <input type='text' onChange={handleChange2} className='textbox'/>
            </div>

            <br></br>
            <div className='sub'>
                {/* 保存ボタン */}
                <button onClick={handleClick} className='button'>保存</button>
            </div>
        </div>
    );
}

export default ExpensesRegister;
