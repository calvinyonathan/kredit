import React, { Component } from 'react'
import { Button, Col, Container, FormControl, FormGroup, FormSelect, Row, Table } from 'react-bootstrap'
import axios from 'axios'
import { API_URL } from '../../const'
export default class DrawdownReport extends Component {
    constructor(props){
        super(props)
        this.state = { customers:[],branch:[] ,company:[]};
    }
    componentDidMount(){
        axios
            .get(API_URL+"/getbranch")
            .then((res) => {
                const branch = res.data.data;
                this.setState({ branch });
              })
              .catch((error) => {
                console.log("Error yaa ", error);
              });

        axios
            .get(API_URL+"/drawdown")
            .then((res) => {
                const customers = res.data.data;
                this.setState({ customers });
              })
              .catch((error) => {
                console.log("Error yaa ", error);
              });
        
        axios
              .get(API_URL+"/getcompany")
              .then((res) => {
                  const company = res.data.data;
                  this.setState({ company });
                })
                .catch((error) => {
                  console.log("Error yaa ", error);
                });
    }
 
    render() {
        let customerList = this.state.customers.map(
            customerList=>(
                <tr>
                    <td>{customerList.RowNumber}</td>
                    <td>{customerList.Ppk}</td>
                    <td>{customerList.Name}</td>
                    <td>{customerList.Channeling_Company}</td>
                    <td>{customerList.DrawdownDate}</td>
                    <td>{customerList.Loan_Amount}</td>
                    <td>{customerList.InterestEffective}%</td>
                    <td><input type={"checkbox"}></input></td>
                </tr>
            )
        )
        let branchList = this.state.branch.map(
            branchList=>(
                <option>{branchList.code}&nbsp;&nbsp;-&nbsp;&nbsp;{branchList.description}</option>
            )
        )
        let companyList = this.state.company.map(
            companyList=>(
                <option>{companyList.company_code}&nbsp;&nbsp;-&nbsp;&nbsp;{companyList.company_short_name}</option>
            )
        )
    return (
        <Container fluid className=''> 
        <Row className='d-flex justify-content-center align-items-center'>
            <Col col='12' >
            <div className='mx-auto content p-5' >
            <h1>Drawdown Report</h1>
            <hr></hr>
            <Row className="d-flex align-items-center justify-content-center">
                <Col className="d-flex align-items-center gap-2 justify-content mb-3">
                    <label>Branch:</label>
                    <FormGroup>  
                        <FormSelect>
                        <option className='d-none'>Please Choose</option>
                        {branchList}
                        </FormSelect>
                    </FormGroup>
                
                    <label>Company:</label>
                    <FormGroup>  
                        <FormSelect>
                            <option className='d-none'>Please Choose</option>
                            {companyList}
                        </FormSelect>
                    </FormGroup>
                   
                    <label>Start: </label>
                <FormGroup>
                    <FormControl type='date'></FormControl>
                </FormGroup>
                <label>End :</label>
                <FormGroup>
                
                    <FormControl type='date'></FormControl>
                </FormGroup>    
                <Button>Submit</Button>   
                </Col>
            </Row>
            <Table striped bordered hover responsive>
                <thead>
                    <tr>
                    <th>No</th>
                    <th>Ppk</th>
                    <th>Name</th>
                    <th>Channeling Company</th>
                    <th>Drawdown Date</th>
                    <th>Loan Amount</th>
                    <th>Interest Eff</th>
                    <th>Action</th>
                    </tr>
                </thead>
                <tbody>
                    {customerList}
                </tbody>
            </Table>
            <Button variant='Primary' style={{ backgroundColor:"#128297",color:"white"}}>
                Approve
            </Button>    
        </div>
            </Col>
        </Row>
    </Container>
       
    )
  }
}
