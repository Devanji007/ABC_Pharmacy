import {  useState } from 'react';
import { toast } from 'react-toastify';
import api from "../api";
import { Card, CardContent, MenuItem, FormControl, Select, TextField, InputLabel} from '@mui/material';
import {  Form,  Row, Col} from 'react-bootstrap';
import LoadingButton from '@mui/lab/LoadingButton';
import CreateItemStyles from '../styles/createItemStyles.module.css';


const CreateItemPage = () => {

    const [itemName, setItemName] = useState('');
    const [unitPrice, setUnitPrice] = useState('');
    const [category, setCategory] = useState('');
  
    const handleSubmit = async (e) => {
      e.preventDefault();
      try {
        const newItem = {
          itemname: itemName,
          unitprice: parseFloat(unitPrice),
          category: category
        };
  
        // Make a POST request to your backend API to create the item
        await api.post('/add_items', newItem);
  
        // Clear form fields after successful submission
        setItemName('');
        setUnitPrice('');
        setCategory('');
        
        // You can add further actions, like displaying a success message or updating the UI
        toast.success('item added');
      } catch (error) {
        toast.error('Error creating item:', error);
        // Handle error, display error message, etc.
      }
    };
  
    const changeCategory = (e) => {
        setCategory(e.target.value);

    }
    return (
      <Form onSubmit={handleSubmit} style={{width:'97%'}}>
         <Row>
            <Col>
                <Card variant="outlined" className={CreateItemStyles.card1}>
                    <CardContent style={{padding:'18px'}}>
                        <h4 style={{margin:0}}>Add Items</h4>
                    </CardContent>
                </Card>
            </Col>
        </Row>
        <Row>
            <Col>
                <Card variant="outlined" className={CreateItemStyles.card} style={{marginBottom:'22px'}}>
                    <CardContent>
                        <Row id={CreateItemStyles.newTicket}>
                        <p><b>New Item</b></p> 
                        </Row>
                    
                        <Row>

                            <Row style={{alignItems:'flex-start', marginTop:'10px'}}>
                                
                                <Col lg={9} xs={6} className='mt-3'><TextField label="Item Name" id="filled-size-normal"   variant="filled" placeholder="Enter Item Name" value={itemName}  onChange={ (e) => setItemName(e.target.value)} required/></Col>
                                
                            </Row>
                            <Row style={{alignItems:'flex-start', marginTop:'10px'}}>
                                <Col lg={9} xs={6} className='mt-3'><TextField label="Unit Price" id="filled-size-normal"   variant="filled" placeholder="Enter Unit Price" value={unitPrice} type='number' step="0.01"  onChange={ (e) => setUnitPrice(e.target.value)} required/></Col>
                               
                            </Row>
                            <Row style={{alignItems:'flex-start', marginTop:'10px'}}>
                              
                                <Col lg={9} xs={6} className='mt-3'>
                                    <FormControl sx={{ m:1, minWidth: 150 }} size="small"> 
                                    <Col lg={3} xs={6}><InputLabel id="demo-simple-select-label" >Category</InputLabel></Col>
                                        <Select value={category} onChange={changeCategory} required  >
                                            <MenuItem value="medications">Medications</MenuItem>
                                            <MenuItem value="petcare">Pet Care</MenuItem>
                                            <MenuItem value="personalcare">Personal Care</MenuItem>
                                            <MenuItem value="firstaid">First Aid</MenuItem>
                                        </Select>
                                    </FormControl>
                                    <br/>
                                   
                                </Col>
                            </Row>
                           
                        </Row>
                            <LoadingButton type="submit"  className='mt-4 mb-4 me-4'  style={{float:'right'}} variant="contained">
                                Create Item
                            </LoadingButton>
                           
                    </CardContent>
                </Card>
            </Col>
        </Row>


</Form>
    
    );


}

export default CreateItemPage;