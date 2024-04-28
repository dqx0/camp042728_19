import './App.css';
import {BrowserRouter as Router,Routes,Route} from "react-router-dom"
import UsingRegister from './UsingRegister'
import ExpensesRegister from './ExpensesRegister';
const App=()=>{
    return(
        <Router>
            <div>
                <Routes>
                    <Route path="/UsingRegister" element={<UsingRegister />} />
                    <Route path="/ExpensesRegister" element={<ExpensesRegister />}/>
                </Routes>
                
            </div>
        </Router>
    )
}
export default App;