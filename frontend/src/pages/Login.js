import React, { Component } from 'react'
import { Button,Image,Col,Row,Container,Card,Form, InputGroup } from "react-bootstrap";
import './Login.css'
import { API_URL,ICONS } from '../const'
import swal from "sweetalert";
import axios from 'axios'
export default class Login extends Component {
    constructor(props) {
		super(props);
		this.state ={
            username : "",
            password : "",  
        }  
	}
    handleUsername = e => {
        const { value } = e.target;
        this.setState({username : value});
    };    
    handlePassword = e => {
        const { value } = e.target;
        this.setState({ password : value});
    };
    handleSubmit = async(event) => {
        event.preventDefault();   
        axios
			.get(API_URL+"/login?nik=" +this.state.username+"&password="+this.state.password)
			.then(res => {
                console.log(res.data.data.len);
                if(res.data.data.length===0)
                {
                    swal({
                        title: "Gagal Login",
                        text: "Username dan Password Salah",
                        icon: "error",
                        button : false,
                        timer : 1500
                    })
                }
                else{
                    localStorage.setItem("nik",this.state.username); 
                    localStorage.setItem("login","true");
                    swal({
                        title: "Sukses Login",
                        text: "Welcome " +this.state.username ,
                        icon: "success",
                        button : false,
                        timer : 1500,
                    }).then(()=>{ 
                    window.location.href="/"   
                    })
                }
            }).catch(error => console.log(error));  

        
	};
    render() {
    return (
        <div className='login-body'>
                        <Container fluid className='login-container mb-3'>
                            <Row className='d-flex justify-content-center align-items-center'>
                                <Col col='12' >
                                    <div className='mx-auto log-shadow p-5' >
                                        <Image src="assets/BankSinarmas.png" width="300" height="200" />
                                        <Card.Body className='w-100 d-flex flex-column'>
                                            <Row>
                                            <InputGroup  className="btn-shadow mb-2" onChange={this.handleUsername} value={this.state.username}>
                                                <InputGroup.Text id="basic-addon1" className='btn-input'><img src={ICONS + "user2.png"} alt={"dd"} style={{ width: "20px", height: "20px", }} /></InputGroup.Text>
                                                <Form.Control
                                                className='btn-input'
                                                placeholder="Username"
                                                aria-label="Username"
                                                aria-describedby="basic-addon1"
                                                />
                                            </InputGroup>
                                            <InputGroup className="btn-shadow mb-3" onChange={this.handlePassword} value={this.state.password} >
                                                <InputGroup.Text className='btn-input' id="basic-addon1"><img src={ICONS + "lock.png"} alt={"dd"} style={{ width: "20px", height: "20px", }} /></InputGroup.Text>
                                                <Form.Control
                                                className='btn-input'
                                                placeholder="Password"
                                                aria-label="Password"
                                                aria-describedby="basic-addon1"
                                                type ="password"
                                                />
                                            </InputGroup>
                                            <Button className='btn-shadow btn-login' variant='Primary' style={{ backgroundColor:"#128297",color:"white"}} onClick={(e)=>this.handleSubmit(e)}>
                                                    Login
                                            </Button>    
                                            </Row>

                                           
                                        </Card.Body>
                                    </div>
                                </Col>
                            </Row>
                        </Container>
                    </div>
                    
    )
  }
}
