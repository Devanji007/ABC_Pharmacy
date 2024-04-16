import { DataGrid } from '@mui/x-data-grid';
import { useEffect, useState } from 'react';

const InvoicePage = () => {

    const [rows, setRows] = useState([]);

    //const getItems = async() => {
      //  try {
        //    let {data} =  await api.get('/items');
           
        //} catch (error) {
          //  toast.error(error);
        //}
      
    //}

    const getInvoices = async() => {
        try {
            let {data} =  await api.get('/invoices');
           
            setRows(data.data)
           
        } catch (error) {
            toast.error(error);
        }
      
    }

    useEffect(() => {
       // getItems();
        getInvoices();
    }, []);


    const columns = [
        {
            field:'invid',
            headerName:'Invoice Id',
            width:150,
            editable: true,
        },

        {
            field:'createdat',
            headerName:'Created at',
            width:150,
            editable: true,
        },
    ]

    const getRowId = (row) => row.invid

    return(
        <>
            <DataGrid columns={columns} rows={rows} getRowId={getRowId} />
        </>
    );
}

export default InvoicePage;