import React from 'react'
import * as AiIcons from 'react-icons/ai'
import * as BsIcons from "react-icons/bs"
import * as MdIcons from 'react-icons/md'
import * as CgIcons from 'react-icons/cg'
export const Sidebarmenu = [
    {
        title: 'Home',
        path: '/',
        icon: <AiIcons.AiFillHome />,
        cName: 'nav-text'
    },
    {
        title: 'Checklist Report',
        path: '/checklist',
        icon: <BsIcons.BsCardChecklist />,
        cName: 'nav-text'
    },
    {
        title: 'Drawdown Report',
        path: '/drawdown',
        icon: <MdIcons.MdPayment />,
        cName: 'nav-text'
    },
    {
        title:'Change Password',
        path:'/changepassword',
        icon: <CgIcons.CgPassword />,
        cName:'nav-text'
    }
]