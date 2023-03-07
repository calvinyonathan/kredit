const customStyles = 
{	header: {		
        style: 
            {			
            minHeight: '56px',		
        },
	},
	headCells: {
		style: {
			'&:not(:last-of-type)': {
				borderRightStyle: 'solid',
				borderRightWidth: '1px',
				borderRightColor: 'black',
			},
		},
	},
	cells: {
		style: {
			'&:not(:last-of-type)': {
				borderRightStyle: 'solid',
				borderRightWidth: '1px',
				borderRightColor: 'black',
			}, 
                
		},
	},
};
export {customStyles}