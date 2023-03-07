import React, { Component } from 'react'
import { Button, Col, Container, Form, FormControl, FormGroup, FormSelect, Row } from 'react-bootstrap'
import axios from 'axios'
import { API_URL } from '../../const'
import swal from 'sweetalert'
import './ChecklistReport.css'
import DataTable from 'react-data-table-component';
import DataTableExtensions from 'react-data-table-component-extensions';
import 'react-data-table-component-extensions/dist/index.css';

import { customStyles } from './ChecklistTable'
export default class ChecklistReport extends Component {
    constructor(props){
        super(props)
        this.state = { customers:[],branch:[] ,company:[],isSubmit:false , 
            currentDate:new Date().toISOString().split('T')[0]
            ,
            checked:[],
            data:[],
        };
        
    }
    componentDidMount(){
        axios
            .get(API_URL+"/getbranch")
            .then((res) => {
                const branch = res.data.data;
                this.setState({ branch });
              })
              .catch((error) => {
                //console.log("Error yaa ", error);
              });
        axios
              .get(API_URL+"/getcompany")
              .then((res) => {
                  const company = res.data.data;
                  this.setState({ company });
                })
                .catch((error) => {
                  //console.log("Error yaa ", error);
                });
    }
    checklist = (ppk,event) => {
            //console.log(ppk)
            if(event.target.checked){
                const checked = [...this.state.checked,{Ppk:ppk}]
                this.setState({checked})
            }
            else{          
                let checkedData = this.state.checked
                checkedData=checkedData.filter((j)=> j.Ppk !== ppk)
                this.setState({checked:checkedData})
            }

    }
    updateApproval = () => {
        //console.log(this.state.checked)
        if(this.state.checked.length === 0){
            swal({
                        title: "Oops Something went wrong   ",
                        text: "Choose Data First !" ,
                        icon: "info",
                        button : false,
                        timer : 1500,
            })
        }
        else{
            axios
            .put('http://localhost:8080/approve', this.state.checked)
            swal({
                title: "Approve Sukses",
                text: "Approve" ,
                icon: "success",
                button : false,
                timer : 1500,
            }).then(()=>{ 
            window.location.href="/checklist"   
            })
        }
     
    }
    handleSubmit = async(event) => {
        event.preventDefault();
        const formData = new FormData(event.currentTarget);
        //console.log(formData.get('branch'))
        if(formData.get('branch')==="Please Choose"){
                swal({
                    title: "Oops Something went wrong   ",
                    text: "Choose Branch First !" ,
                    icon: "info",
                    button : false,
                    timer : 1000,
            })
        }
        else if(formData.get('company')==="Please Choose")
        {
            swal({
                title: "Oops Something went wrong   ",
                text: "Choose Company First !" ,
                icon: "info",
                button : false,
                timer : 1000,
            })
        }
        else{
        this.setState({isSubmit:true})
        axios
            .get(API_URL+"/checklistBranch?branch="+formData.get('branch')+"&company="+formData.get('company')+"&startdate="+formData.get('startDate')+"&enddate="+formData.get('endDate'))
            .then((res) => {
                const customers = res.data.data;
                this.setState({ customers });
              })
              .catch((error) => {
                //onsole.log("Error yaa ", error);
              });
        }
    }
    render() {
        let data=[]
        this.state.customers.map(
            (customerList,id)=>(
                data = [...data,{
                    id : id+1,
                    ppk : customerList.Ppk,
                    name : customerList.Name,
                    company:customerList.Channeling_Company,
                    drawdowndate:customerList.DrawdownDate,
                    loanamount:customerList.Loan_Amount,
                    interesteff:customerList.InterestEffective,
                    action :<input type={"checkbox"} onChange={(e) => this.checklist(customerList.Ppk,e)}></input>
                }]
                
            )
        )
        let branchList = this.state.branch.map(
            branchList=>(
                <option value={branchList.code}>{branchList.code}&nbsp;&nbsp;-&nbsp;&nbsp;{branchList.description}</option>
            )
        )
        let companyList = this.state.company.map(
            companyList=>(
                <option value={companyList.company_short_name}>{companyList.company_code}&nbsp;&nbsp;-&nbsp;&nbsp;{companyList.company_short_name}</option>
            )
        )
    const columns = [
            {
                name: 'PPK',
                selector: 'ppk',
                sortable: true,
                maxWidth: '200px'
            },
            {
                name: 'Name',
                selector: 'name',
                sortable: true,
            },
            {
                name: 'Channeling Company',
                selector: 'company',
                sortable: true,
            },
            {
                name: 'Drawdown Date',
                selector: 'drawdowndate',
                sortable: true,
            },
            {
                name: 'Loan Amount',
                selector: 'loanamount',
                sortable: true,
            },
            {
                name: 'Interest Eff',
                selector:'interesteff',
                sortable: true,
                maxWidth: '100 px'
            },
            {
             
                cell : d=> d.action,
                maxWidth: '10px'
            }
    ];
    const data2 =[
        {   
            drawdowndate:'Please Choose Data',
        }
    ]
    const data3 =[
        {   
            drawdowndate:'No Data Found',
        }
    ]
    const tableData = {
        columns,
        data,
    };   
  
    
 
    return (
        <Container fluid className=''> 
        <Row className='d-flex justify-content-center align-items-center'>
            <Col col='12' >
            <div className='mx-auto content p-5' >
            <h1>Checklist Report</h1>
            <hr></hr>
            <Form onSubmit={(e)=>this.handleSubmit(e)}>            
                <Row className="d-flex align-items-center justify-content-center">
                    <Col className="d-flex align-items-center gap-2 justify-content mb-3">
                        <label>Branch:</label>
                        <FormGroup>  
                            <FormSelect name='branch'>
                            <option className='d-none'>Please Choose</option>
                            {branchList}
                            </FormSelect>
                        </FormGroup>
                    
                        <label>Company:</label>
                        <FormGroup>  
                            <FormSelect name='company'>
                                <option className='d-none'>Please Choose</option>
                                {companyList}
                            </FormSelect>
                        </FormGroup>
                    
                    <label>Start: </label>
                    <FormGroup>
                        <FormControl type='date' name='startDate' defaultValue={this.state.currentDate}></FormControl>
                    </FormGroup>
                    <label>End :</label>
                    <FormGroup>
                    
                        <FormControl type='date' name='endDate' defaultValue={this.state.currentDate}></FormControl>
                    </FormGroup>    
                    <Button type='submit'>Submit</Button>   
                    </Col>
                </Row>
            </Form>
            {this.state.isSubmit===false  ?
            <DataTable
            columns={columns}
            data={data2}
            pagination
            responsive
            /> 
            :
                this.state.customers.length===0 ?
                    <DataTable
                    columns={columns}
                    data={data3}
                    pagination
                    responsive
                    /> 
                        :
                    <DataTableExtensions
                        {...tableData}
                    >
                    <DataTable
                        columns={columns}
                        data={data}
                        pagination
                        responsive
                        customStyles={customStyles}
                        defaultSortField="id"
                    defaultSortAsc={false}
                    />
                    </DataTableExtensions>   
             }  
            
            <Button onClick={() => this.updateApproval()} variant='Primary' style={{ backgroundColor:"#128297",color:"white"}}>
                Approve
            </Button>    
        </div>
            </Col>
        </Row>
    </Container>
       
    )
  }
}
