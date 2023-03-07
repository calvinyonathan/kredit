import React, { Component } from 'react'
import { Image,Col,Row,Container} from "react-bootstrap";

import './home.css'
export default class Home extends Component {
  render() {
    return (
      <div>
    <Container fluid className=''>
      <Row className='d-flex justify-content-center align-items-center  '>
          <Col col='12' >
            <div className='mx-auto content p-5 text-center home ' >
              <hr></hr>
              <h1>Welcome Program Bantu Kredit</h1>
              <h4>Supported By :</h4>
              <Image src="assets/BankSinarmas.png" width="300" height="200" />
              <hr></hr>
            </div>
          </Col>
      </Row>
  </Container>
  
  </div>
    )
  }
}
