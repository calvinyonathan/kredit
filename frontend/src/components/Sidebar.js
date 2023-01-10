import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import { ICONS } from '../const';
import './Sidebar.css';
export default function Sidebar() {
  const logout = () => {
    localStorage.removeItem('name');
    window.location.href="/"   
  }
    return (
      <div className="d-flex">
      <div>
        <Nav className="flex-column" id="nav-trigger">
          <div className="d-flex align-items-center menubar-brand">
            {/* <img id="image-trigger"onClick={() => collapseNavbar()} src={Logo} alt={"Logo Action Figure "}></img> */}
            <div className="d-flex flex-column">
              <h3 className="">Program Bantu Kredit</h3>
              <h6 className="">Hello, Calvin</h6>
            </div>
          </div>
          <hr></hr>
          <div className="">
              <Nav.Link 
                href="/checklist">
                  <img src={ICONS + "user2.png"} alt={"dd"} style={{ width: "20px", height: "20px",marginRight:"10px"}} />
                  <span className="">Checklist Report</span>
              </Nav.Link>
              <Nav.Link> 
              <img src={ICONS + "user2.png"} alt={"dd"} style={{ width: "20px", height: "20px", marginRight:"10px" }} />
                  <span className="">Drawdown Report</span>
              </Nav.Link>
              <hr></hr>
          </div>
          <div className="">
            <div className="d-flex">
              <div className="user-summary">
              <Nav.Link onClick={() => logout()}> 
                  <img src={ICONS + "log-out.png"} alt={"dd"} style={{ width: "20px", height: "20px", marginRight:"10px"}} />
                  <span className="" >Log Out</span>
              </Nav.Link>
              </div>
            </div>
          </div>
        </Nav>
      </div>
      </div>
    )
}