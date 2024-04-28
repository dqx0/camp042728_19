import React from "react";
import { Link } from "react-router-dom"

export const Header = () => {
    return(
        <div>
            <header className="header">おかねみえるくん</header>
            <br></br>
            <div className="sub" >
                <Link to="/MonthlyMoneyInput">１ヶ月の使用可能金額を登録する</Link>
            </div>
            <br></br>
            <div className="sub">
                <Link to="/UsingRegister">使用できる金額を見る</Link>
            </div>
            <br></br>
            <div className="sub">
                <Link to="/ExpensesRegister">使った金額を登録する</Link>
            </div>
        </div>
    );
}

export default Header