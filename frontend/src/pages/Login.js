import React, { Component } from 'react'
import { Button,Image,Col,Row,Container,Card,Form, InputGroup } from "react-bootstrap";
import './Login.css'
import { ICONS } from '../const'
export default class Login extends Component {
  render() {
    return (
        <div className='login-body'>
                        <Container fluid className='login-container mb-3'>
                            <Row className='d-flex justify-content-center align-items-center'>
                                <Col col='12' >
                                    <div className='my-5 mx-auto log-shadow text-center p-5' >
                                        <Image src="assets/pega-logo.svg" width="250" height="50" />
                                        <Card.Body className='w-100 d-flex flex-column'>
                                            <Row>

                                            <InputGroup  className="btn-shadow mb-2 mt-5" >
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
                                                placeholder="Username"
                                                aria-label="Username"
                                                aria-describedby="basic-addon1"
                                                type ="password"
                                                />
                                            </InputGroup>
                                            <Button className='btn-shadow btn-input' variant='Primary' style={{ backgroundColor:"#128297",color:"white"}} onClick={(e)=>this.handleSubmit(e)}>
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
