import React from 'react'
import * as FaIcons from 'react-icons/fa'
import * as AiIcons from 'react-icons/ai'
import * as IoIcons from 'react-icons/io'
import * as BsIcons from "react-icons/bs"
import * as MdIcons from 'react-icons/md'
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
    }
]