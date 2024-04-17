import {  useState } from 'react';
import { toast } from 'react-toastify';
import api from "../api";
import { Card, CardContent, MenuItem, FormControl, Select, InputLabel} from '@mui/material';
import { Container, Form, Button, Row, Col} from 'react-bootstrap';
import LoadingButton from '@mui/lab/LoadingButton';
import CreateItemStyles from '../styles/createItemStyles.module.css';
import TextField from '@mui/material/TextField';


const CreateInvoicePage = () => {

    const [cusName, setCusName] = useState('');
    const [mobNo, setMobNo] = useState('');
    const [email, setEmail] = useState('');
    const [address, setAddress] = useState('');
    const [billingType, setBillingType] = useState('');
    const [totalamount, setTotalAmount] = useState('');
  
    const handleSubmit = async (e) => {
      e.preventDefault();
      try {
        const newInvoice = {
            cusname: cusName,
            mobno: mobNo,
            email: email,
            address: address,
            billingtype: billingType,
            totalamount: totalamount
          };
  
        // Make a POST request to your backend API to create the item
        await api.post('/create_invoice', newInvoice);
  
        // Clear form fields after successful submission
        setCusName('');
        setMobNo('');
        setEmail('');
        setAddress('');
        setBillingType('');
        setTotalAmount('');

        // You can add further actions, like displaying a success message or updating the UI
        toast.success('invoice created');
      } catch (error) {
        toast.error('Error creating invoice:', error);
        // Handle error, display error message, etc.
      }
    };

    const changeType = (e) => {
        setBillingType(e.target.value);

    }
  
    return (
      <Form onSubmit={handleSubmit} style={{width:'97%'}}>
         <Row>
            <Col>
                <Card variant="outlined" className={CreateItemStyles.card1}>
                    <CardContent style={{padding:'18px'}}>
                        <h4 style={{margin:0}}>Create Invoice</h4>
                    </CardContent>
                </Card>
            </Col>
        </Row>
        <Row>
            <Col>
                <Card variant="outlined" className={CreateItemStyles.card} style={{marginBottom:'22px'}}>
                    <CardContent>
                        <Row id={CreateItemStyles.newTicket}>
                        <p><b>New Invoice</b></p> 
                        </Row>
                    
                        <Row>

                            <Row style={{alignItems:'flex-start', marginTop:'10px'}}>
                            
                               {/* <Col lg={3} xs={6}><label htmlFor="name" className={CreateItemStyles.lbl}>Customer Name<span className={CreateItemStyles.require}><b>*</b></span></label></Col>*/}
                                <Col lg={9} xs={6} className='mt-3'><TextField label="Customer Name" id="filled-size-normal"   variant="filled" placeholder='Enter Customer Name' value={cusName}  onChange={ (e) => setCusName(e.target.value)} required/></Col>
                            </Row>
                            <Row style={{alignItems:'flex-start', marginTop:'10px'}}>
                                <Col lg={9} xs={6} className='mt-3'><TextField label="Mob No" id="filled-size-normal"   variant="filled" placeholder='Enter Mob No' value={mobNo}  onChange={ (e) => setMobNo(e.target.value)} required/></Col>
                            </Row>
                            <Row style={{alignItems:'flex-start', marginTop:'10px'}}>
                                <Col lg={9} xs={6} className='mt-3'><TextField label="Email" id="filled-size-normal"   variant="filled" placeholder='Enter Email' value={email} type='email' onChange={ (e) => setEmail(e.target.value)} required/></Col>
                                
                            </Row>
                            <Row style={{alignItems:'flex-start', marginTop:'10px'}}>                        
                                <Col lg={9} xs={6} className='mt-3'><TextField label="Address" id="filled-size-normal"   variant="filled" placeholder='Enter Address' value={address} onChange={ (e) => setAddress(e.target.value)} required/></Col>
                            </Row>
                            <Row style={{alignItems:'flex-start', marginTop:'20px'}}>
                                
                                <Col lg={9} xs={6} className='mt-3'>
                                    <FormControl sx={{ m:1, minWidth: 150 }} size="small" > 
                                    <Col lg={3} xs={6}><InputLabel id="demo-simple-select-label" >Billing Type</InputLabel></Col>
                                        <Select value={billingType} onChange={changeType} required  >
                                            <MenuItem value="cash">Cash</MenuItem>
                                            <MenuItem value="card">Card</MenuItem>
                                        </Select>
                                    </FormControl>
                                    <br/>
                                   
                                </Col>
                            </Row>

                            <Row style={{alignItems:'flex-start', marginTop:'10px'}}>
                                <Col lg={9} xs={6} className='mt-3'><TextField label="TotalAmount" id="filled-size-normal"   variant="filled"  value={totalamount} type='number'  step="0.01"  onChange={ (e) => setTotalAmount(e.target.value)} required/></Col>
                            </Row>
                           
                        </Row>
                            <LoadingButton type="submit"  className='mt-4 mb-4 me-4'  style={{float:'right', marginBottom:'20px'}} variant="contained">
                                Create Invoice
                            </LoadingButton>
                           
                    </CardContent>
                </Card>
            </Col>
        </Row>


</Form>
    
    );


}

export default CreateInvoicePage;