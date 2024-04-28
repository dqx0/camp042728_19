import './App.css';
import {BrowserRouter as Router,Routes,Route} from "react-router-dom"
import UsingRegister from './UsingRegister'
import ExpensesRegister from './ExpensesRegister';
import Header from './Header';
import MonthlyMoneyInput from './MonthlyMoneyInput';
const App=()=>{
    return(
        <Router>
            <div>
                <Header />
                <Routes>
                    <Route path="/UsingRegister" element={<UsingRegister />} />
                    <Route path="/ExpensesRegister" element={<ExpensesRegister />}/>
                    <Route path='MonthlyMoneyInput' element={<MonthlyMoneyInput/>}/>
                </Routes>
                
            </div>
        </Router>
    )
}
export default App;