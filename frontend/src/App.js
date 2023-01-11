import logo from './logo.svg';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import './App.css';
import Login from './pages/Login';
import Sidebars from './components/Sidebar';
import Home from './pages/Home/home';
import ChecklistReport from './components/ChecklistReport/ChecklistReport';
import { Col, Container, Row } from 'react-bootstrap';
import DrawdownReport from './components/DrawdownReport/DrawdownReport';

function App() {
  if(localStorage.getItem("name")=="Calvin"){
    return(
      <BrowserRouter>
      <div className='d-flex'>
           <div className=''>
               <Sidebars></Sidebars>
           </div>
                    <div className='page'>
                        <Routes>
                          <Route path="/" element={<Home />} exact />
                          <Route path="/checklist" element={<ChecklistReport />} exact />
                          <Route path="/drawdown" element={<DrawdownReport />} exact />
                        </Routes>
                    </div> 

      </div>
     </BrowserRouter>
    )
  }
  else{
  return (
       <BrowserRouter>
              <Routes>
                <Route path="/" element={<Login />} exact />
             
              </Routes>
      </BrowserRouter>
  );
  }
}

export default App;
