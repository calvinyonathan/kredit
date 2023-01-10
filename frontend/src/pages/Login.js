import React, { Component } from 'react'
import { Button,Image,Col,Row,Container,Card,Form, InputGroup } from "react-bootstrap";
import './Login.css'
import { ICONS } from '../const'
import swal from "sweetalert";
export default class Login extends Component {
    handleSubmit = async(event) => {
        localStorage.setItem("name","Calvin"); 
        swal({
            title: "Sukses Login",
            text: "Welcome Calvin" ,
            icon: "success",
            button : false,
            timer : 1500,
        }).then(()=>{ 
        window.location.href="/"   
        })
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

                                            <InputGroup  className="btn-shadow mb-2" >
                                                <InputGroup.Text id="basic-addon1" className='btn-input'><img src={ICONS + "user2.png"} alt={"dd"} style={{ width: "20px", height: "20px", }} /></InputGroup.Text>
                                                <Form.Control
                                                className='btn-input'
                                                placeholder="Username"
                                                aria-label="Username"
                                                aria-describedby="basic-addon1"
                                                />
                                            </InputGroup>
                                            <InputGroup className="btn-shadow mb-3" >
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
