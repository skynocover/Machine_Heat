package main

import (
	"github.com/zserge/lorca"
	"log"
	"net/url"
	"io/ioutil"
	"strconv"
	"fmt"
)

var (
	viewtitle      = "視窗標題"
	inputlabel     = "標題"
	bottonname     = "按鈕名稱"
	vieweight  int = 900
	viewheight int = 850

	ui, err = lorca.New("", "", vieweight, viewheight)

	opera topera
	motor tmotor
	pump tpump
	worker tworker
	coolant tcoolant
	moving Tmoving
)

var html = `<html>
	<head>
		<title>Heat of Machine</title>
		
		<style type="text/css">
			.tabs{      
				position:relative;      
				width:100%;      
				height:100%;  
			}

			.tab{ float:left;}

			.tab > input[type=radio]{      
				display:none;  
			} 

			.tab > label{      
				display:block;      
				position:relative;      
				min-width:40px;      
				height:20px;      
				margin-right:-1px;      
				padding:5px 15px;     
				border:1px solid #AAA;  
			} 

			.tab > .content{      
				display:none;      
				position:absolute;      
				top:30px;      
				right:0;      
				bottom:0;      
				left:0;      
				z-index:1;      
				padding:10px;     
			}    

			.tab > input[type=radio]:checked + label{     
				border-bottom:1px solid transparent;      
				z-index:2;  
			}

			.tab > input[type=radio]:checked ~ .content{      
				display:block;      
			}

			/*
			Solid State by HTML5 UP
			html5up.net | @ajlkn
			Free for personal and commercial use under the CCA 3.0 license (html5up.net/license)
			*/

			/* Reset */

			html, body, div, span, applet, object, iframe, h1, h2, h3, h4, h5, h6, p, blockquote, pre, a, abbr, acronym, address, big, cite, code, del, dfn, em, img, ins, kbd, q, s, samp, small, strike, strong, sub, sup, tt, var, b, u, i, center, dl, dt, dd, ol, ul, li, fieldset, form, label, legend, table, caption, tbody, tfoot, thead, tr, th, td, article, aside, canvas, details, embed, figure, figcaption, footer, header, hgroup, menu, nav, output, ruby, section, summary, time, mark, audio, video {
				margin: 0;
				padding: 0;
				border: 0;
				font-size: 100%;
				font: inherit;
				vertical-align: baseline;
			}

			article, aside, details, figcaption, figure, footer, header, hgroup, menu, nav, section {
				display: block;
			}

			body {
				line-height: 1;
			}

			ol, ul {
				list-style: none;
			}

			blockquote, q {
				quotes: none;
			}

			blockquote:before, blockquote:after, q:before, q:after {
				content: '';
				content: none;
			}

			table {
				border-collapse: collapse;
				border-spacing: 0;
			}

			body {
				-webkit-text-size-adjust: none;
			}



			/* Grid */

			.row {
				border-bottom: solid 1px transparent;
				-moz-box-sizing: border-box;
				-webkit-box-sizing: border-box;
				box-sizing: border-box;
			}

			.row > * {
				float: left;
				-moz-box-sizing: border-box;
				-webkit-box-sizing: border-box;
				box-sizing: border-box;
			}

			.row:after, .row:before {
				content: '';
				display: block;
				clear: both;
				height: 0;
			}

			.row.uniform > * > :first-child {
				margin-top: 0;
			}

			.row.uniform > * > :last-child {
				margin-bottom: 0;
			}

			.row.\30 \25 > * {
				padding: 0 0 0 0em;
			}

			.row.\30 \25 {
				margin: 0 0 -1px 0em;
			}

			.row.uniform.\30 \25 > * {
				padding: 0em 0 0 0em;
			}

			.row.uniform.\30 \25 {
				margin: 0em 0 -1px 0em;
			}

			.row > * {
				padding: 0 0 0 1.75em;
			}

			.row {
				margin: 0 0 -1px -1.75em;
			}

			.row.uniform > * {
				padding: 1.75em 0 0 1.75em;
			}

			.row.uniform {
				margin: -1.75em 0 -1px -1.75em;
			}

			.row.\32 00\25 > * {
				padding: 0 0 0 3.5em;
			}

			.row.\32 00\25 {
				margin: 0 0 -1px -3.5em;
			}

			.row.uniform.\32 00\25 > * {
				padding: 3.5em 0 0 3.5em;
			}

			.row.uniform.\32 00\25 {
				margin: -3.5em 0 -1px -3.5em;
			}

			.row.\31 50\25 > * {
				padding: 0 0 0 2.625em;
			}

			.row.\31 50\25 {
				margin: 0 0 -1px -2.625em;
			}

			.row.uniform.\31 50\25 > * {
				padding: 2.625em 0 0 2.625em;
			}

			.row.uniform.\31 50\25 {
				margin: -2.625em 0 -1px -2.625em;
			}

			.row.\35 0\25 > * {
				padding: 0 0 0 0.875em;
			}

			.row.\35 0\25 {
				margin: 0 0 -1px -0.875em;
			}

			.row.uniform.\35 0\25 > * {
				padding: 0.875em 0 0 0.875em;
			}

			.row.uniform.\35 0\25 {
				margin: -0.875em 0 -1px -0.875em;
			}

			.row.\32 5\25 > * {
				padding: 0 0 0 0.4375em;
			}

			.row.\32 5\25 {
				margin: 0 0 -1px -0.4375em;
			}

			.row.uniform.\32 5\25 > * {
				padding: 0.4375em 0 0 0.4375em;
			}

			.row.uniform.\32 5\25 {
				margin: -0.4375em 0 -1px -0.4375em;
			}

			.\31 2u, .\31 2u\24 {
				width: 100%;
				clear: none;
				margin-left: 0;
			}

			.\31 1u, .\31 1u\24 {
				width: 91.6666666667%;
				clear: none;
				margin-left: 0;
			}

			.\31 0u, .\31 0u\24 {
				width: 83.3333333333%;
				clear: none;
				margin-left: 0;
			}

			.\39 u, .\39 u\24 {
				width: 75%;
				clear: none;
				margin-left: 0;
			}

			.\38 u, .\38 u\24 {
				width: 66.6666666667%;
				clear: none;
				margin-left: 0;
			}

			.\37 u, .\37 u\24 {
				width: 58.3333333333%;
				clear: none;
				margin-left: 0;
			}

			.\36 u, .\36 u\24 {
				width: 50%;
				clear: none;
				margin-left: 0;
			}

			.\35 u, .\35 u\24 {
				width: 41.6666666667%;
				clear: none;
				margin-left: 0;
			}

			.\34 u, .\34 u\24 {
				width: 33.3333333333%;
				clear: none;
				margin-left: 0;
			}

			.\33 u, .\33 u\24 {
				width: 33.3333%;
				clear: none;
				margin-left: 0;
			}

			.\32 u, .\32 u\24 {
				width: 16.6666666667%;
				clear: none;
				margin-left: 0;
			}

			.\31 u, .\31 u\24 {
				width: 8.3333333333%;
				clear: none;
				margin-left: 0;
			}

			.\31 2u\24 + *,
			.\31 1u\24 + *,
			.\31 0u\24 + *,
			.\39 u\24 + *,
			.\38 u\24 + *,
			.\37 u\24 + *,
			.\36 u\24 + *,
			.\35 u\24 + *,
			.\34 u\24 + *,
			.\33 u\24 + *,
			.\32 u\24 + *,
			.\31 u\24 + * {
				clear: left;
			}

			.\-11u {
				margin-left: 91.66667%;
			}

			.\-10u {
				margin-left: 83.33333%;
			}

			.\-9u {
				margin-left: 75%;
			}

			.\-8u {
				margin-left: 66.66667%;
			}

			.\-7u {
				margin-left: 58.33333%;
			}

			.\-6u {
				margin-left: 50%;
			}

			.\-5u {
				margin-left: 41.66667%;
			}

			.\-4u {
				margin-left: 33.33333%;
			}

			.\-3u {
				margin-left: 33.3333%;
			}

			.\-2u {
				margin-left: 16.66667%;
			}

			.\-1u {
				margin-left: 8.33333%;
			}

			@media screen and (max-width: 1680px) {

				.row > * {
					padding: 0 0 0 1.75em;
				}

				.row {
					margin: 0 0 -1px -1.75em;
				}

				.row.uniform > * {
					padding: 1.75em 0 0 1.75em;
				}

				.row.uniform {
					margin: -1.75em 0 -1px -1.75em;
				}

				.row.\32 00\25 > * {
					padding: 0 0 0 3.5em;
				}

				.row.\32 00\25 {
					margin: 0 0 -1px -3.5em;
				}

				.row.uniform.\32 00\25 > * {
					padding: 3.5em 0 0 3.5em;
				}

				.row.uniform.\32 00\25 {
					margin: -3.5em 0 -1px -3.5em;
				}

				.row.\31 50\25 > * {
					padding: 0 0 0 2.625em;
				}

				.row.\31 50\25 {
					margin: 0 0 -1px -2.625em;
				}

				.row.uniform.\31 50\25 > * {
					padding: 2.625em 0 0 2.625em;
				}

				.row.uniform.\31 50\25 {
					margin: -2.625em 0 -1px -2.625em;
				}

				.row.\35 0\25 > * {
					padding: 0 0 0 0.875em;
				}

				.row.\35 0\25 {
					margin: 0 0 -1px -0.875em;
				}

				.row.uniform.\35 0\25 > * {
					padding: 0.875em 0 0 0.875em;
				}

				.row.uniform.\35 0\25 {
					margin: -0.875em 0 -1px -0.875em;
				}

				.row.\32 5\25 > * {
					padding: 0 0 0 0.4375em;
				}

				.row.\32 5\25 {
					margin: 0 0 -1px -0.4375em;
				}

				.row.uniform.\32 5\25 > * {
					padding: 0.4375em 0 0 0.4375em;
				}

				.row.uniform.\32 5\25 {
					margin: -0.4375em 0 -1px -0.4375em;
				}

				.\31 2u\28xlarge\29, .\31 2u\24\28xlarge\29 {
					width: 100%;
					clear: none;
					margin-left: 0;
				}

				.\31 1u\28xlarge\29, .\31 1u\24\28xlarge\29 {
					width: 91.6666666667%;
					clear: none;
					margin-left: 0;
				}

				.\31 0u\28xlarge\29, .\31 0u\24\28xlarge\29 {
					width: 83.3333333333%;
					clear: none;
					margin-left: 0;
				}

				.\39 u\28xlarge\29, .\39 u\24\28xlarge\29 {
					width: 75%;
					clear: none;
					margin-left: 0;
				}

				.\38 u\28xlarge\29, .\38 u\24\28xlarge\29 {
					width: 66.6666666667%;
					clear: none;
					margin-left: 0;
				}

				.\37 u\28xlarge\29, .\37 u\24\28xlarge\29 {
					width: 58.3333333333%;
					clear: none;
					margin-left: 0;
				}

				.\36 u\28xlarge\29, .\36 u\24\28xlarge\29 {
					width: 50%;
					clear: none;
					margin-left: 0;
				}

				.\35 u\28xlarge\29, .\35 u\24\28xlarge\29 {
					width: 41.6666666667%;
					clear: none;
					margin-left: 0;
				}

				.\34 u\28xlarge\29, .\34 u\24\28xlarge\29 {
					width: 33.3333333333%;
					clear: none;
					margin-left: 0;
				}

				.\33 u\28xlarge\29, .\33 u\24\28xlarge\29 {
					width: 33.3333%;
					clear: none;
					margin-left: 0;
				}

				.\32 u\28xlarge\29, .\32 u\24\28xlarge\29 {
					width: 16.6666666667%;
					clear: none;
					margin-left: 0;
				}

				.\31 u\28xlarge\29, .\31 u\24\28xlarge\29 {
					width: 8.3333333333%;
					clear: none;
					margin-left: 0;
				}

				.\31 2u\24\28xlarge\29 + *,
				.\31 1u\24\28xlarge\29 + *,
				.\31 0u\24\28xlarge\29 + *,
				.\39 u\24\28xlarge\29 + *,
				.\38 u\24\28xlarge\29 + *,
				.\37 u\24\28xlarge\29 + *,
				.\36 u\24\28xlarge\29 + *,
				.\35 u\24\28xlarge\29 + *,
				.\34 u\24\28xlarge\29 + *,
				.\33 u\24\28xlarge\29 + *,
				.\32 u\24\28xlarge\29 + *,
				.\31 u\24\28xlarge\29 + * {
					clear: left;
				}

				.\-11u\28xlarge\29 {
					margin-left: 91.66667%;
				}

				.\-10u\28xlarge\29 {
					margin-left: 83.33333%;
				}

				.\-9u\28xlarge\29 {
					margin-left: 75%;
				}

				.\-8u\28xlarge\29 {
					margin-left: 66.66667%;
				}

				.\-7u\28xlarge\29 {
					margin-left: 58.33333%;
				}

				.\-6u\28xlarge\29 {
					margin-left: 50%;
				}

				.\-5u\28xlarge\29 {
					margin-left: 41.66667%;
				}

				.\-4u\28xlarge\29 {
					margin-left: 33.33333%;
				}

				.\-3u\28xlarge\29 {
					margin-left: 33.3333%;
				}

				.\-2u\28xlarge\29 {
					margin-left: 16.66667%;
				}

				.\-1u\28xlarge\29 {
					margin-left: 8.33333%;
				}

			}

			@media screen and (max-width: 1280px) {

				.row > * {
					padding: 0 0 0 1.75em;
				}

				.row {
					margin: 0 0 -1px -1.75em;
				}

				.row.uniform > * {
					padding: 1.75em 0 0 1.75em;
				}

				.row.uniform {
					margin: -1.75em 0 -1px -1.75em;
				}

				.row.\32 00\25 > * {
					padding: 0 0 0 3.5em;
				}

				.row.\32 00\25 {
					margin: 0 0 -1px -3.5em;
				}

				.row.uniform.\32 00\25 > * {
					padding: 3.5em 0 0 3.5em;
				}

				.row.uniform.\32 00\25 {
					margin: -3.5em 0 -1px -3.5em;
				}

				.row.\31 50\25 > * {
					padding: 0 0 0 2.625em;
				}

				.row.\31 50\25 {
					margin: 0 0 -1px -2.625em;
				}

				.row.uniform.\31 50\25 > * {
					padding: 2.625em 0 0 2.625em;
				}

				.row.uniform.\31 50\25 {
					margin: -2.625em 0 -1px -2.625em;
				}

				.row.\35 0\25 > * {
					padding: 0 0 0 0.875em;
				}

				.row.\35 0\25 {
					margin: 0 0 -1px -0.875em;
				}

				.row.uniform.\35 0\25 > * {
					padding: 0.875em 0 0 0.875em;
				}

				.row.uniform.\35 0\25 {
					margin: -0.875em 0 -1px -0.875em;
				}

				.row.\32 5\25 > * {
					padding: 0 0 0 0.4375em;
				}

				.row.\32 5\25 {
					margin: 0 0 -1px -0.4375em;
				}

				.row.uniform.\32 5\25 > * {
					padding: 0.4375em 0 0 0.4375em;
				}

				.row.uniform.\32 5\25 {
					margin: -0.4375em 0 -1px -0.4375em;
				}

				.\31 2u\28large\29, .\31 2u\24\28large\29 {
					width: 100%;
					clear: none;
					margin-left: 0;
				}

				.\31 1u\28large\29, .\31 1u\24\28large\29 {
					width: 91.6666666667%;
					clear: none;
					margin-left: 0;
				}

				.\31 0u\28large\29, .\31 0u\24\28large\29 {
					width: 83.3333333333%;
					clear: none;
					margin-left: 0;
				}

				.\39 u\28large\29, .\39 u\24\28large\29 {
					width: 75%;
					clear: none;
					margin-left: 0;
				}

				.\38 u\28large\29, .\38 u\24\28large\29 {
					width: 66.6666666667%;
					clear: none;
					margin-left: 0;
				}

				.\37 u\28large\29, .\37 u\24\28large\29 {
					width: 58.3333333333%;
					clear: none;
					margin-left: 0;
				}

				.\36 u\28large\29, .\36 u\24\28large\29 {
					width: 50%;
					clear: none;
					margin-left: 0;
				}

				.\35 u\28large\29, .\35 u\24\28large\29 {
					width: 41.6666666667%;
					clear: none;
					margin-left: 0;
				}

				.\34 u\28large\29, .\34 u\24\28large\29 {
					width: 33.3333333333%;
					clear: none;
					margin-left: 0;
				}

				.\33 u\28large\29, .\33 u\24\28large\29 {
					width: 33.3333%;
					clear: none;
					margin-left: 0;
				}

				.\32 u\28large\29, .\32 u\24\28large\29 {
					width: 16.6666666667%;
					clear: none;
					margin-left: 0;
				}

				.\31 u\28large\29, .\31 u\24\28large\29 {
					width: 8.3333333333%;
					clear: none;
					margin-left: 0;
				}

				.\31 2u\24\28large\29 + *,
				.\31 1u\24\28large\29 + *,
				.\31 0u\24\28large\29 + *,
				.\39 u\24\28large\29 + *,
				.\38 u\24\28large\29 + *,
				.\37 u\24\28large\29 + *,
				.\36 u\24\28large\29 + *,
				.\35 u\24\28large\29 + *,
				.\34 u\24\28large\29 + *,
				.\33 u\24\28large\29 + *,
				.\32 u\24\28large\29 + *,
				.\31 u\24\28large\29 + * {
					clear: left;
				}

				.\-11u\28large\29 {
					margin-left: 91.66667%;
				}

				.\-10u\28large\29 {
					margin-left: 83.33333%;
				}

				.\-9u\28large\29 {
					margin-left: 75%;
				}

				.\-8u\28large\29 {
					margin-left: 66.66667%;
				}

				.\-7u\28large\29 {
					margin-left: 58.33333%;
				}

				.\-6u\28large\29 {
					margin-left: 50%;
				}

				.\-5u\28large\29 {
					margin-left: 41.66667%;
				}

				.\-4u\28large\29 {
					margin-left: 33.33333%;
				}

				.\-3u\28large\29 {
					margin-left: 33.3333%;
				}

				.\-2u\28large\29 {
					margin-left: 16.66667%;
				}

				.\-1u\28large\29 {
					margin-left: 8.33333%;
				}

			}

			@media screen and (max-width: 980px) {

				.row > * {
					padding: 0 0 0 1.75em;
				}

				.row {
					margin: 0 0 -1px -1.75em;
				}

				.row.uniform > * {
					padding: 1.75em 0 0 1.75em;
				}

				.row.uniform {
					margin: -1.75em 0 -1px -1.75em;
				}

				.row.\32 00\25 > * {
					padding: 0 0 0 3.5em;
				}

				.row.\32 00\25 {
					margin: 0 0 -1px -3.5em;
				}

				.row.uniform.\32 00\25 > * {
					padding: 3.5em 0 0 3.5em;
				}

				.row.uniform.\32 00\25 {
					margin: -3.5em 0 -1px -3.5em;
				}

				.row.\31 50\25 > * {
					padding: 0 0 0 2.625em;
				}

				.row.\31 50\25 {
					margin: 0 0 -1px -2.625em;
				}

				.row.uniform.\31 50\25 > * {
					padding: 2.625em 0 0 2.625em;
				}

				.row.uniform.\31 50\25 {
					margin: -2.625em 0 -1px -2.625em;
				}

				.row.\35 0\25 > * {
					padding: 0 0 0 0.875em;
				}

				.row.\35 0\25 {
					margin: 0 0 -1px -0.875em;
				}

				.row.uniform.\35 0\25 > * {
					padding: 0.875em 0 0 0.875em;
				}

				.row.uniform.\35 0\25 {
					margin: -0.875em 0 -1px -0.875em;
				}

				.row.\32 5\25 > * {
					padding: 0 0 0 0.4375em;
				}

				.row.\32 5\25 {
					margin: 0 0 -1px -0.4375em;
				}

				.row.uniform.\32 5\25 > * {
					padding: 0.4375em 0 0 0.4375em;
				}

				.row.uniform.\32 5\25 {
					margin: -0.4375em 0 -1px -0.4375em;
				}

				.\31 2u\28medium\29, .\31 2u\24\28medium\29 {
					width: 100%;
					clear: none;
					margin-left: 0;
				}

				.\31 1u\28medium\29, .\31 1u\24\28medium\29 {
					width: 91.6666666667%;
					clear: none;
					margin-left: 0;
				}

				.\31 0u\28medium\29, .\31 0u\24\28medium\29 {
					width: 83.3333333333%;
					clear: none;
					margin-left: 0;
				}

				.\39 u\28medium\29, .\39 u\24\28medium\29 {
					width: 75%;
					clear: none;
					margin-left: 0;
				}

				.\38 u\28medium\29, .\38 u\24\28medium\29 {
					width: 66.6666666667%;
					clear: none;
					margin-left: 0;
				}

				.\37 u\28medium\29, .\37 u\24\28medium\29 {
					width: 58.3333333333%;
					clear: none;
					margin-left: 0;
				}

				.\36 u\28medium\29, .\36 u\24\28medium\29 {
					width: 50%;
					clear: none;
					margin-left: 0;
				}

				.\35 u\28medium\29, .\35 u\24\28medium\29 {
					width: 41.6666666667%;
					clear: none;
					margin-left: 0;
				}

				.\34 u\28medium\29, .\34 u\24\28medium\29 {
					width: 33.3333333333%;
					clear: none;
					margin-left: 0;
				}

				.\33 u\28medium\29, .\33 u\24\28medium\29 {
					width: 33.3333%;
					clear: none;
					margin-left: 0;
				}

				.\32 u\28medium\29, .\32 u\24\28medium\29 {
					width: 16.6666666667%;
					clear: none;
					margin-left: 0;
				}

				.\31 u\28medium\29, .\31 u\24\28medium\29 {
					width: 8.3333333333%;
					clear: none;
					margin-left: 0;
				}

				.\31 2u\24\28medium\29 + *,
				.\31 1u\24\28medium\29 + *,
				.\31 0u\24\28medium\29 + *,
				.\39 u\24\28medium\29 + *,
				.\38 u\24\28medium\29 + *,
				.\37 u\24\28medium\29 + *,
				.\36 u\24\28medium\29 + *,
				.\35 u\24\28medium\29 + *,
				.\34 u\24\28medium\29 + *,
				.\33 u\24\28medium\29 + *,
				.\32 u\24\28medium\29 + *,
				.\31 u\24\28medium\29 + * {
					clear: left;
				}

				.\-11u\28medium\29 {
					margin-left: 91.66667%;
				}

				.\-10u\28medium\29 {
					margin-left: 83.33333%;
				}

				.\-9u\28medium\29 {
					margin-left: 75%;
				}

				.\-8u\28medium\29 {
					margin-left: 66.66667%;
				}

				.\-7u\28medium\29 {
					margin-left: 58.33333%;
				}

				.\-6u\28medium\29 {
					margin-left: 50%;
				}

				.\-5u\28medium\29 {
					margin-left: 41.66667%;
				}

				.\-4u\28medium\29 {
					margin-left: 33.33333%;
				}

				.\-3u\28medium\29 {
					margin-left: 33.3333%;
				}

				.\-2u\28medium\29 {
					margin-left: 16.66667%;
				}

				.\-1u\28medium\29 {
					margin-left: 8.33333%;
				}

			}






			/* Basic */

			@-ms-viewport {
				width: device-width;
			}

			body {
				-ms-overflow-style: scrollbar;
			}



			body {
				background-color: #2e3141;
				background-image: linear-gradient(to top, rgba(46, 49, 65, 0.8), rgba(46, 49, 65, 0.8)), url("../../images/bg.jpg");
				background-size: auto, cover;
				background-attachment: fixed, fixed;
				background-position: center, center;
			}

				body.is-loading *, body.is-loading *:before, body.is-loading *:after {
					-moz-animation: none !important;
					-webkit-animation: none !important;
					-ms-animation: none !important;
					animation: none !important;
					-moz-transition: none !important;
					-webkit-transition: none !important;
					-ms-transition: none !important;
					transition: none !important;
				}

			/* Type */

			body, input, select, textarea {
				color: #ffffff;
				font-family: "Source Sans Pro", Helvetica, sans-serif;
				font-size: 16.5pt;
				font-weight: 300;
				line-height: 1.65;
			}

				@media screen and (max-width: 1680px) {

					body, input, select, textarea {
						font-size: 13pt;
					}

				}

				@media screen and (max-width: 1280px) {

					body, input, select, textarea {
						font-size: 12pt;
					}

				}

				@media screen and (max-width: 980px) {

					body, input, select, textarea {
						font-size: 12pt;
					}

				}




			a {
				-moz-transition: color 0.2s ease-in-out, border-bottom-color 0.2s ease-in-out;
				-webkit-transition: color 0.2s ease-in-out, border-bottom-color 0.2s ease-in-out;
				-ms-transition: color 0.2s ease-in-out, border-bottom-color 0.2s ease-in-out;
				transition: color 0.2s ease-in-out, border-bottom-color 0.2s ease-in-out;
				border-bottom: dotted 1px rgba(255, 255, 255, 0.35);
				color: #ffffff;
				text-decoration: none;
			}

				a:hover {
					border-bottom-color: transparent;
					color: #ffffff !important;
				}

				a.special:not(.button) {
					text-decoration: none;
					border-bottom: 0;
					display: block;
					font-family: Raleway, Helvetica, sans-serif;
					font-size: 0.8em;
					font-weight: 700;
					letter-spacing: 0.1em;
					margin: 0 0 2em 0;
					text-transform: uppercase;
				}

					a.special:not(.button):before {
						-moz-osx-font-smoothing: grayscale;
						-webkit-font-smoothing: antialiased;
						font-family: FontAwesome;
						font-style: normal;
						font-weight: normal;
						text-transform: none !important;
					}

					a.special:not(.button):before {
						-moz-transition: background-color 0.2s ease-in-out;
						-webkit-transition: background-color 0.2s ease-in-out;
						-ms-transition: background-color 0.2s ease-in-out;
						transition: background-color 0.2s ease-in-out;
						border-radius: 100%;
						border: solid 2px rgba(255, 255, 255, 0.125);
						content: '\f105';
						display: inline-block;
						font-size: 1.25em;
						height: 2em;
						line-height: 1.65em;
						margin-right: 0.85em;
						text-align: center;
						text-indent: 0.15em;
						vertical-align: middle;
						width: 2em;
					}

					a.special:not(.button):hover:before {
						background-color: rgba(255, 255, 255, 0.025);
					}

					a.special:not(.button):active:before {
						background-color: rgba(255, 255, 255, 0.075);
					}

			strong, b {
				color: #ffffff;
				font-weight: 600;
			}

			em, i {
				font-style: italic;
			}

			p {
				margin: 0 0 2em 0;
			}

			h1, h2, h3, h4, h5, h6 {
				color: #ffffff;
				font-family: Raleway, Helvetica, sans-serif;
				font-weight: 700;
				letter-spacing: 0.1em;
				margin: 0 0 1em 0;
				text-transform: uppercase;
			}

				h1 a, h2 a, h3 a, h4 a, h5 a, h6 a {
					color: inherit;
					text-decoration: none;
				}

				h1 span, h2 span, h3 span, h4 span, h5 span, h6 span {
					font-weight: 200;
				}

				h1.major, h2.major, h3.major, h4.major, h5.major, h6.major {
					padding-bottom: 1em;
					border-bottom: solid 2px rgba(255, 255, 255, 0.125);
				}

			h2 {
				font-size: 1.2em;
			}

			h3 {
				font-size: 0.9em;
			}

			h4 {
				font-size: 0.7em;
			}

			h5 {
				font-size: 0.7em;
			}

			h6 {
				font-size: 0.7em;
			}

			sub {
				font-size: 0.8em;
				position: relative;
				top: 0.5em;
			}

			sup {
				font-size: 0.8em;
				position: relative;
				top: -0.5em;
			}

			blockquote {
				border-left: solid 4px rgba(255, 255, 255, 0.125);
				font-style: italic;
				margin: 0 0 2em 0;
				padding: 0.5em 0 0.5em 2em;
			}

			code {
				background: rgba(255, 255, 255, 0.025);
				border-radius: 5px;
				border: solid 2px rgba(255, 255, 255, 0.125);
				font-family: "Courier New", monospace;
				font-size: 0.9em;
				margin: 0 0.25em;
				padding: 0.25em 0.65em;
			}

			pre {
				-webkit-overflow-scrolling: touch;
				font-family: "Courier New", monospace;
				font-size: 0.9em;
				margin: 0 0 2em 0;
			}

				pre code {
					display: block;
					line-height: 1.75em;
					padding: 1em 1.5em;
					overflow-x: auto;
				}

			hr {
				border: 0;
				border-bottom: solid 2px rgba(255, 255, 255, 0.125);
				margin: 2.5em 0;
			}

				hr.major {
					margin: 4em 0;
				}

			.align-left {
				text-align: left;
			}

			.align-center {
				text-align: center;
			}

			.align-right {
				text-align: right;
			}



			/* Form */

			form {
				margin: 0 0 2em 0;
			}

				form .field {
					margin: 0 0 2em 0;
				}

					form .field label {
						margin-top: -1.5em;
					}

					form .field:first-child label {
						margin-top: 0;
					}

				form > :last-child {
					margin-bottom: 0;
				}

			label {
				color: #ffffff;
				display: block;
				font-family: Raleway, Helvetica, sans-serif;
				font-size: 0.8em;
				font-weight: 700;
				letter-spacing: 0.1em;
				margin: 0 0 0.7em 0;
				text-transform: uppercase;
			}

			input[type="text"],
			input[type="password"],
			input[type="email"],
			input[type="tel"],
			select,
			textarea {
				-moz-appearance: none;
				-webkit-appearance: none;
				-ms-appearance: none;
				appearance: none;
				background: rgba(255, 255, 255, 0.025);
				border-radius: 5px;
				border: none;
				border: solid 2px rgba(255, 255, 255, 0.125);
				color: inherit;
				display: block;
				outline: 0;
				padding: 0 1em;
				text-decoration: none;
				width: 100%;
			}

				input[type="text"]:invalid,
				input[type="password"]:invalid,
				input[type="email"]:invalid,
				input[type="tel"]:invalid,
				select:invalid,
				textarea:invalid {
					box-shadow: none;
				}

				input[type="text"]:focus,
				input[type="password"]:focus,
				input[type="email"]:focus,
				input[type="tel"]:focus,
				select:focus,
				textarea:focus {
					border-color: #5b6ba6;
				}

			.select-wrapper {
				text-decoration: none;
				display: block;
				position: relative;
			}

				.select-wrapper:before {
					-moz-osx-font-smoothing: grayscale;
					-webkit-font-smoothing: antialiased;
					font-family: FontAwesome;
					font-style: normal;
					font-weight: normal;
					text-transform: none !important;
				}

				.select-wrapper:before {
					color: rgba(255, 255, 255, 0.125);
					content: '\f078';
					display: block;
					height: 2.75em;
					line-height: 2.75em;
					pointer-events: none;
					position: absolute;
					right: 0;
					text-align: center;
					top: 0;
					width: 2.75em;
				}

				.select-wrapper select::-ms-expand {
					display: none;
				}

			select option {
				color: #ffffff;
				background: #2e3141;
			}

			input[type="text"],
			input[type="password"],
			input[type="email"],
			select {
				height: 2.75em;
			}

			textarea {
				padding: 0.75em 1em;
			}

			input[type="checkbox"],
			input[type="radio"] {
				-moz-appearance: none;
				-webkit-appearance: none;
				-ms-appearance: none;
				appearance: none;
				display: block;
				float: left;
				margin-right: -2em;
				opacity: 0;
				width: 1em;
				z-index: -1;
			}

				input[type="checkbox"] + label,
				input[type="radio"] + label {
					text-decoration: none;
					color: #ffffff;
					cursor: pointer;
					display: inline-block;
					font-size: 1em;
					font-family: "Source Sans Pro", Helvetica, sans-serif;
					text-transform: none;
					letter-spacing: 0;
					font-weight: 300;
					padding-left: 2.4em;
					padding-right: 0.75em;
					position: relative;
				}

					input[type="checkbox"] + label:before,
					input[type="radio"] + label:before {
						-moz-osx-font-smoothing: grayscale;
						-webkit-font-smoothing: antialiased;
						font-family: FontAwesome;
						font-style: normal;
						font-weight: normal;
						text-transform: none !important;
					}

					input[type="checkbox"] + label:before,
					input[type="radio"] + label:before {
						background: rgba(255, 255, 255, 0.025);
						border-radius: 5px;
						border: solid 2px rgba(255, 255, 255, 0.125);
						content: '';
						display: inline-block;
						height: 1.65em;
						left: 0;
						line-height: 1.58125em;
						position: absolute;
						text-align: center;
						top: 0;
						width: 1.65em;
					}

				input[type="checkbox"]:checked + label:before,
				input[type="radio"]:checked + label:before {
					background: #ffffff;
					border-color: #ffffff;
					content: '\f00c';
					color: #2e3141;
				}

				input[type="checkbox"]:focus + label:before,
				input[type="radio"]:focus + label:before {
					border-color: #4c5c96;
				}

			input[type="checkbox"] + label:before {
				border-radius: 5px;
			}

			input[type="radio"] + label:before {
				border-radius: 100%;
			}

			::-webkit-input-placeholder {
				color: rgba(255, 255, 255, 0.35) !important;
				opacity: 1.0;
			}

			:-moz-placeholder {
				color: rgba(255, 255, 255, 0.35) !important;
				opacity: 1.0;
			}

			::-moz-placeholder {
				color: rgba(255, 255, 255, 0.35) !important;
				opacity: 1.0;
			}

			:-ms-input-placeholder {
				color: rgba(255, 255, 255, 0.35) !important;
				opacity: 1.0;
			}

			.polyfill-placeholder {
				color: rgba(255, 255, 255, 0.35) !important;
				opacity: 1.0;
			}

			/* Box */

			.box {
				border-radius: 5px;
				border: solid 2px rgba(255, 255, 255, 0.125);
				margin-bottom: 2em;
				padding: 1.5em;
			}

				.box > :last-child,
				.box > :last-child > :last-child,
				.box > :last-child > :last-child > :last-child {
					margin-bottom: 0;
				}

				.box.alt {
					border: 0;
					border-radius: 0;
					padding: 0;
				}



			/* Button */

			input[type="submit"],
			input[type="reset"],
			input[type="button"],
			button,
			.button {
				-moz-appearance: none;
				-webkit-appearance: none;
				-ms-appearance: none;
				appearance: none;
				-moz-transition: background-color 0.2s ease-in-out;
				-webkit-transition: background-color 0.2s ease-in-out;
				-ms-transition: background-color 0.2s ease-in-out;
				transition: background-color 0.2s ease-in-out;
				background-color: transparent;
				border-radius: 5px;
				border: 0;
				box-shadow: inset 0 0 0 2px rgba(255, 255, 255, 0.125);
				color: #ffffff !important;
				cursor: pointer;
				display: inline-block;
				font-family: Raleway, Helvetica, sans-serif;
				font-size: 0.8em;
				font-weight: 700;
				height: 3.75em;
				letter-spacing: 0.1em;
				line-height: 3.75em;
				padding: 0 2.25em;
				text-align: center;
				text-decoration: none;
				text-transform: uppercase;
				white-space: nowrap;
			}

				input[type="submit"]:hover,
				input[type="reset"]:hover,
				input[type="button"]:hover,
				button:hover,
				.button:hover {
					background-color: rgba(255, 255, 255, 0.025);
				}

				input[type="submit"]:active,
				input[type="reset"]:active,
				input[type="button"]:active,
				button:active,
				.button:active {
					background-color: rgba(255, 255, 255, 0.075);
				}

				input[type="submit"].icon:before,
				input[type="reset"].icon:before,
				input[type="button"].icon:before,
				button.icon:before,
				.button.icon:before {
					margin-right: 0.5em;
					color: rgba(255, 255, 255, 0.35);
				}

				input[type="submit"].special,
				input[type="reset"].special,
				input[type="button"].special,
				button.special,
				.button.special {
					background-color: #4c5c96;
					box-shadow: none;
				}

					input[type="submit"].special:hover,
					input[type="reset"].special:hover,
					input[type="button"].special:hover,
					button.special:hover,
					.button.special:hover {
						background-color: #53639e;
					}

					input[type="submit"].special:active,
					input[type="reset"].special:active,
					input[type="button"].special:active,
					button.special:active,
					.button.special:active {
						background-color: #45558d;
					}

					input[type="submit"].special.icon:before,
					input[type="reset"].special.icon:before,
					input[type="button"].special.icon:before,
					button.special.icon:before,
					.button.special.icon:before {
						color: #7985b0;
					}

				input[type="submit"].fit,
				input[type="reset"].fit,
				input[type="button"].fit,
				button.fit,
				.button.fit {
					display: block;
					margin: 0 0 1em 0;
					width: 100%;
				}

				input[type="submit"].small,
				input[type="reset"].small,
				input[type="button"].small,
				button.small,
				.button.small {
					font-size: 0.6em;
				}

				input[type="submit"].big,
				input[type="reset"].big,
				input[type="button"].big,
				button.big,
				.button.big {
					font-size: 1em;
				}

				input[type="submit"].disabled, input[type="submit"]:disabled,
				input[type="reset"].disabled,
				input[type="reset"]:disabled,
				input[type="button"].disabled,
				input[type="button"]:disabled,
				button.disabled,
				button:disabled,
				.button.disabled,
				.button:disabled {
					opacity: 0.25;
				}
				
			</style>
		<meta charset="utf-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1" />
		<script>
			function GWtype(){
				if (document.getElementById('gwType').value==0) {
					document.getElementById('gwR').value = ""
					document.getElementById('gwR').disabled = true;
				}else{
					document.getElementById('gwR').disabled = false;
				}
			}
			function targetChange(){
				var targetype  = document.getElementById("targeType"); 
				if (document.getElementById('target').value==0) {
					targetype.options.length =0
					targetype.options.add(new Option("內藏主軸",0));
					targetype.options.add(new Option("非內藏主軸",1));
				}else{
					targetype.options.length =0
					targetype.options.add( new Option("不含馬達介面冷卻",0));
					targetype.options.add(new Option("含馬達介面冷卻",1));
				}
			}
			function targetTotal() {
				var cond = document.getElementById("target").value+document.getElementById("targeType").value
				switch (cond) {
					case "00":
						document.getElementById('targetResult').innerHTML="計算結果:軸承+馬達發熱"
						if (document.getElementById('w').value=="") {
							alert("請先計算軸承發熱")
						}else if (document.getElementById('motorW').value==""){
							alert("請先計算馬達發熱")
						}else{
							document.getElementById('targetTotal').value=(document.getElementById('w').value/1000+document.getElementById('motorW').value/1).toFixed(4)
						}
						break;
					case "01":
						document.getElementById('targetResult').innerHTML="計算結果:軸承發熱"
						if (document.getElementById('w').value=="") {
							alert("請先計算軸承發熱")
						}else{
							document.getElementById('targetTotal').value=(document.getElementById('w').value/1000).toFixed(4)
						}
						break;

					case "10":
						document.getElementById('targetResult').innerHTML="計算結果:軸承+螺桿發熱"
						if (document.getElementById('w').value=="") {
							alert("請先計算軸承發熱")
						}else if (document.getElementById('bsW').value==""){
							alert("請先計算螺桿發熱")
						}else{
							document.getElementById('targetTotal').value=(document.getElementById('w').value/1000+document.getElementById('bsW').value/1000).toFixed(4)
						}
						break;
					
					default:
						document.getElementById('targetResult').innerHTML="計算結果:軸承+螺桿+馬達發熱"
						if (document.getElementById('w').value=="") {
							alert("請先計算軸承發熱")
						}else if (document.getElementById('bsW').value==""){
							alert("請先計算螺桿發熱")
						}else if(document.getElementById('motorW').value==""){
							alert("請先計算馬達發熱")
						}else{
							document.getElementById('targetTotal').value=(document.getElementById('w').value/1000+document.getElementById('bsW').value/1000).toFixed(4)
						}
						break;
				}
				document.getElementById("targetW").value=(document.getElementById("targetTotal").value*1.3*1000).toFixed(4)
				document.getElementById("targetCal").value=(document.getElementById("targetTotal").value*1.3*1000*0.86).toFixed(4)
			}
		
		</script>
	</head>
	<body>			
		<div class="tabs">        
			<div class="tab">          
				<input type="radio" id="tab1" name="tabs" checked="checked">          
				<label for="tab1">軸承發熱</label>          
				<div class="content">
					<section>
						<h3 class="major">軸承參數</h3>
						<form method="post" action="#">
							<div class="row uniform">
								<div class="6u 12u$(xsmall)" style="width:33.3333%">
									<label for="demo-name">Dm值(mm)</label>
									<input type="text" name="demo-name" id="dm" value="140" />
								</div>
								<div class="6u 12u$(xsmall)" style="width:33.3333%">
									<label for="demo-name">靜額定負荷(N)</label>
									<input type="text" name="demo-name" id="c0" value="105000" />
								</div>
								<div class="12u$" style="width:33.3333%">
									<label for="demo-category">軸承型式</label>
									<div class="select-wrapper">
										<select name="demo-category" id="bearingtype">
											<option value="0">滾珠軸承</option>
											<option value="1">滾柱軸承</option>
										</select>
									</div>
								</div>
								<div class="6u 12u$(xsmall)" style="width:33.3333%">
									<label for="demo-name">軸承顆數</label>
									<input type="text" name="demo-name" id="pcs" value="2" />
								</div>
								
								<div class="6u 12u$(xsmall)" style="width:33.3333%">
									<label for="demo-name">接觸角(度)</label>
									<input type="text" name="demo-name" id="angle" value="15" />
								</div>
								
								<div class="12u$" style="width:33.3333%">
									<label for="demo-category">列數</label>
									<div class="select-wrapper">
										<select name="demo-category" id="row">
											<option value="1">單列</option>
											<option value="2">複列</option>
										</select>
									</div>
								</div>
								
							</div>
							
						</form>
						<h3 class="major">運轉參數</h3>
						<form method="post" action="#">
							<div class="row uniform">
								<div class="6u 12u$(xsmall)" style="width:33.3333%">
									<label for="demo-name">轉速(rpm)</label>
									<input type="text" name="demo-name" id="rpm" value="6000" />
								</div>
								<div class="6u 12u$(xsmall)" style="width:33.3333%">
									<label for="demo-name">推力荷重(kgf)</label>
									<input type="text" name="demo-name" id="fu" value="8.6" />
								</div>
								<div class="6u 12u$(xsmall)" style="width:33.3333%">
									<label for="demo-name">徑向荷重(kgf)</label>
									<input type="text" name="demo-name" id="fr" value="6.8" />
								</div>
								<div class="6u 12u$(xsmall)" style="width:33.3333%">
									<label for="demo-name">黏度(cst)</label>
									<input type="text" name="demo-name" id="v" value="25" />
								</div>
								<div class="12u$" style="width:33.3333%">
									<label for="demo-category">潤滑方式</label>
									<div class="select-wrapper">
										<select name="demo-category" id="lubetype">
											<option value="0">油氣潤滑</option>
											<option value="1">油脂潤滑</option>
											<option value="2">噴射潤滑</option>
										</select>
									</div>
								</div>
								
							</div>
						</form>
						<div class="12u$">
							<ul class="actions">
								<li><input type="submit" value="Calc" class="special" onclick="execute(document.querySelector('#dm').value
																									,document.querySelector('#pcs').value
																									,document.querySelector('#v').value
																									,document.querySelector('#c0').value
																									,document.querySelector('#angle').value
																									,document.querySelector('#bearingtype').value
																									,document.querySelector('#rpm').value
																									,document.querySelector('#fu').value
																									,document.querySelector('#fr').value
																									,document.querySelector('#lubetype').value
																									,document.querySelector('#row').value)" /></li>
							</ul>
						</div>
					</form>
						<h3 class="major">計算結果</h3>
						<form method="post" action="#">
							<div class="row uniform">
								<div class="6u 12u$(xsmall)" style="width:25%">
									<label for="demo-name">速度項(kgf*m)</label>
									<output type="text" name="demo-name" id="mv" value=""></output>
								</div>
								<div class="6u 12u$(xsmall)" style="width:25%">
									<label for="demo-name">荷重項(kgf*m)</label>
									<output type="text" name="demo-name" id="ml" value="" ></output>
								</div>
								<div class="6u 12u$(xsmall)" style="width:25%">
									<label for="demo-name">發熱量(kcal/h)</label>
									<output type="text" name="demo-name" id="q" value="" ></output>
								</div>
								<div class="6u 12u$(xsmall)" style="width:25%">
									<label for="demo-name">發熱量(w)</label>
									<output type="text" name="demo-name" id="w" value="" ></output>
								</div>								
							</div>
						</form>
						
						
					</section>
				</div>      
			</div>            
			<div class="tab">          
				<input type="radio" id="tab2" name="tabs">          
				<label for="tab2">馬達發熱</label>          
				<div class="content">
					<h3 class="major">馬達參數</h3>
					<form method="post" action="#">
						<div class="row uniform">	
							<div class="6u 12u$(xsmall)" style="width:33.3333%">
								<label for="demo-category">馬達類型</label>
									<div class="select-wrapper">
									<select name="demo-category" id="motorType">
										<option value="0">交流</option>
										<option value="1">直流</option>
										<option value="2">伺服</option>
									</select>
								</div>
							</div>							
							<div class="6u 12u$(xsmall)" style="width:33.3333%">
								<label for="demo-name">馬達出力(KW)</label>
								<input type="text" name="demo-name" id="output" value="25" />
							</div>
						</div>
					</form>
					<div class="12u$">
						<ul class="actions">
							<li><input type="submit" value="Calc" class="special" onclick="motorHeat(document.querySelector('#motorType').value
																								,document.querySelector('#output').value)" /></li>
						</ul>
					</div>
					<h3 class="major">計算結果</h3>
					<form method="post" action="#">
						<div class="row uniform">							
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-name">發熱量(Kw)</label>
								<output type="text" name="demo-name" id="motorW" value="" ></output>
							</div>								
						</div>
					</form>
				</div>      
			</div>            
			<div class="tab">          
				<input type="radio" id="tab3" name="tabs">          
				<label for="tab3">加工發熱</label>          
				<div class="content">
					<h3 class="major">加工參數</h3>
					<form method="post" action="#">
						<div class="row uniform">													
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-name">主軸馬達出力(KW)</label>
								<input type="text" name="demo-name" id="spindlePower" value="12" />
							</div>
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-name">粗加工時間比例(0~1)</label>
								<input type="text" name="demo-name" id="roughTime" value="0.7" />
							</div>
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-name">精加工時間比例(0~1)</label>
								<input type="text" name="demo-name" id="finishTime" value="0.2" />
							</div>
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-category">有無排屑機</label>
									<div class="select-wrapper">
									<select name="demo-category" id="chipremoval">
										<option value="1">無排屑機</option>
										<option value="0.8">有排屑機</option>
									</select>
								</div>
							</div>	
						</div>
					</form>
					<h3 class="major">幫浦發熱</h3>
					<form method="post" action="#">
						<div class="row uniform">													
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-name">幫浦1馬達出力(KW)</label>
								<input type="text" name="demo-name" id="power1" value="12" />
							</div>
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-name">幫浦1馬達數量</label>
								<input type="text" name="demo-name" id="num1" value="3" />
							</div>
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-name">幫浦2馬達出力(KW)</label>
								<input type="text" name="demo-name" id="power2" value="15" />
							</div>
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-name">幫浦2馬達數量</label>
								<input type="text" name="demo-name" id="num2" value="1" />
							</div>
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-name">幫浦3馬達出力(KW)</label>
								<input type="text" name="demo-name" id="power3" value="0" />
							</div>
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-name">幫浦3馬達數量</label>
								<input type="text" name="demo-name" id="num3" value="0" />
							</div>
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-name">幫浦4馬達出力(KW)</label>
								<input type="text" name="demo-name" id="power4" value="0" />
							</div>
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-name">幫浦4馬達數量</label>
								<input type="text" name="demo-name" id="num4" value="0" />
							</div>
							
						</div>
					</form>
					<h3 class="major">切削液表面熱</h3>
					<form method="post" action="#">
						<div class="row uniform">
							<div class="6u 12u$(xsmall)" style="width:33%">
								<label for="demo-name">水箱表面積(m^2)</label>
								<input type="text" name="demo-name" id="surface" value="5" />
							</div>
							<div class="6u 12u$(xsmall)" style="width:33%">
								<label for="demo-name">周圍溫度與設定溫度差(°C)</label>
								<input type="text" name="demo-name" id="temperature" value="2" />
							</div>
							
						</div>
					</form>
					<div class="12u$">
						<ul class="actions">
							
							<li><input type="submit" value="Calc" class="special" onclick="coolantHeat(document.querySelector('#spindlePower').value
																									,document.querySelector('#roughTime').value
																										,document.querySelector('#finishTime').value
																										,document.querySelector('#chipremoval').value
																										,document.querySelector('#power1').value
																										,document.querySelector('#num1').value
																										,document.querySelector('#power2').value
																										,document.querySelector('#num2').value
																										,document.querySelector('#power3').value
																										,document.querySelector('#num3').value
																										,document.querySelector('#power4').value
																										,document.querySelector('#num4').value
																										,document.querySelector('#surface').value
																										,document.querySelector('#temperature').value)" /></li>
						
						</ul>
					</div>
					<h3 class="major">計算結果</h3>
					<form method="post" action="#">
						<div class="row uniform">							
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-name">加工熱(kw)</label>
								<output type="text" name="demo-name" id="workerW" value="" ></output>
							</div>	
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-name">幫浦熱(kw)</label>
								<output type="text" name="demo-name" id="pumpW" value="" ></output>
							</div>
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-name">切削液熱(kw)</label>
								<output type="text" name="demo-name" id="coolantW" value="" ></output>
							</div>		
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-name">總發熱(kw)</label>
								<output type="text" name="demo-name" id="workW" value="" ></output>
							</div>						
						</div>
					</form>
				</div>      
			</div>    
			<div class="tab">          
					<input type="radio" id="tab4" name="tabs">          
					<label for="tab4">驅動發熱</label>          
					<div class="content">
						<h3 class="major">驅動參數</h3>
						<form method="post" action="#">
							<div class="row uniform">													
								<div class="6u 12u$(xsmall)" style="width:25%">
									<label for="demo-name">工作物重量(Kg)</label>
									<input type="text" name="demo-name" id="weight" value="300" />
								</div>
								<div class="6u 12u$(xsmall)" style="width:25%">
									<label for="demo-name">軸向負荷(kg)</label>
									<input type="text" name="demo-name" id="axiaload" value="7" />
								</div>
								<div class="6u 12u$(xsmall)" style="width:25%">
									<label for="demo-name">移動速度(mm/min)</label>
									<input type="text" name="demo-name" id="movingv" value="14000" />
								</div>									
							</div>
						</form>
						<h3 class="major">螺桿參數</h3>
						<form method="post" action="#">
							<div class="row uniform">													
								<div class="6u 12u$(xsmall)" style="width:25%">
									<label for="demo-name">導程(mm)</label>
									<input type="text" name="demo-name" id="bsLead" value="10" />
								</div>
								<div class="6u 12u$(xsmall)" style="width:25%">
									<label for="demo-name">效率</label>
									<input type="text" name="demo-name" id="bsEff" value="0.9" />
								</div>
								<div class="6u 12u$(xsmall)" style="width:25%">
									<label for="demo-name">轉速(rpm)</label>
									<output type="text" name="demo-name" id="bsRpm" value="" ></output>
								</div>
								<div class="6u 12u$(xsmall)" style="width:25%">
									<label for="demo-name">扭矩(kgf*mm)</label>
									<output type="text" name="demo-name" id="bsTorque" value="" ></output>
								</div>
							</div>
						</form>		
						<h3 class="major">導軌參數</h3>
						<form method="post" action="#">
							<div class="row uniform">	
								<div class="6u 12u$(xsmall)" style="width:25%">
									<label for="demo-category">導軌型式</label>
									<div class="select-wrapper">
										<select name="demo-category" id="gwType" onchange="GWtype()">
											<option value="0">滑動</option>
											<option value="1" selected="true">滾動</option>
										</select>
									</div>
								</div>												
								<div class="6u 12u$(xsmall)" style="width:25%">
									<label for="demo-name">滾動體半徑(mm)</label>
									<input type="text" name="demo-name" id="gwR" value="3.175" />
								</div>								
								<div class="6u 12u$(xsmall)" style="width:25%">
									<label for="demo-name">摩擦係數</label>
									<input type="text" name="demo-name" id="gwU" value="0.02" />
								</div>								
							</div>							
						</form>				
						<div class="12u$">
							<ul class="actions">
								
								<li><input type="submit" value="Calc" class="special" onclick="movingHeat(document.querySelector('#weight').value
																										,document.querySelector('#axiaload').value
																										,document.querySelector('#movingv').value
																										,document.querySelector('#bsLead').value
																										,document.querySelector('#bsEff').value
																										,document.querySelector('#gwR').value
																										,document.querySelector('#gwU').value
																										,document.querySelector('#gwType').value)" /></li>
							
							</ul>
						</div>
						<h3 class="major">計算結果</h3>
						<form method="post" action="#">
							<div class="row uniform">							
								<div class="6u 12u$(xsmall)" style="width:25%">
									<label for="demo-name">螺桿發熱(w)</label>
									<output type="text" name="demo-name" id="bsW" value="" ></output>
								</div>	
								<div class="6u 12u$(xsmall)" style="width:25%">
									<label for="demo-name">線軌發熱(w)</label>
									<output type="text" name="demo-name" id="gwW" value="" ></output>
								</div>														
							</div>
						</form>
					</div>      
			</div> 
			<div class="tab">          
				<input type="radio" id="tab5" name="tabs">          
				<label for="tab5">計算範本</label>          
				<div class="content">
					<h3 class="major">計算目標</h3>
					<form method="post" action="#">
						<div class="row uniform">	
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-category">目標部件</label>
								<div class="select-wrapper">
									<select name="demo-category" id="target" onchange="targetChange()">
										<option value="0">主軸</option>
										<option value="1">螺桿</option>
									</select>
								</div>
							</div>		
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-category">部件型式</label>
								<div class="select-wrapper">
									<select name="demo-category" id="targeType" >
										<option value="0">內藏主軸</option>
										<option value="1">非內藏主軸</option>
									</select>
								</div>
							</div>
						</div>							
					</form>				
					<div class="12u$">
						<ul class="actions">
							
							<li><input type="submit" value="Calc" class="special" onclick="targetTotal()" /></li>
						
						</ul>
					</div>
					<h3 class="major" id="targetResult">計算結果:軸承+馬達發熱</h3>
					<form method="post" action="#">
						<div class="row uniform">
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-name">總發熱(kw)</label>
								<output type="text" name="demo-name" id="targetTotal" value="" ></output>
							</div>		
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-name">冷凍機建議值(w)</label>
								<output type="text" name="demo-name" id="targetW" value="" ></output>
							</div>
							<div class="6u 12u$(xsmall)" style="width:25%">
								<label for="demo-name">冷凍機建議值(kcal/h)</label>
								<output type="text" name="demo-name" id="targetCal" value="" ></output>
							</div>											
						</div>
					</form>
				</div>      
		</div> 
		</div>    
	</body>
	</html> `


func main() {
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	//ui.Load("data:text/html," + url.PathEscape(read("file/home.html")))
	ui.Load("data:text/html," + url.PathEscape(html))
	ui.Bind("execute", func(dm string,pcs string,v string,c0 string,angle string,bearingtype string,rpm string,fu string,fr string,lubetype string,row string) {
		opera.Dm,_=strconv.ParseFloat(dm, 64)
		opera.Pcs,_=strconv.ParseFloat(pcs, 64)
		opera.V,_=strconv.ParseFloat(v, 64)
		opera.C0,_=strconv.ParseFloat(c0, 64)
		opera.Angle,_=strconv.ParseFloat(angle, 64)
		opera.Bearingtype,_=strconv.ParseFloat(bearingtype, 64)
		opera.Rpm,_=strconv.ParseFloat(rpm, 64)
		opera.Fu,_=strconv.ParseFloat(fu, 64)
		opera.Fr,_=strconv.ParseFloat(fr, 64)
		opera.Lubetype,_=strconv.ParseFloat(lubetype, 64)
		opera.Row,_=strconv.ParseFloat(row, 64)
		opera.calc()
		opera.show()
		fmt.Println(opera)
	})
	ui.Bind("motorHeat", func(motorType string,output string) {
		motor.Type,_=strconv.ParseFloat(motorType, 64)
		motor.Power,_=strconv.ParseFloat(output, 64)
		ui.Eval(`document.getElementById("motorW").innerText = `+strconv.FormatFloat(motor.q(), 'E', 6, 64))		
		fmt.Println(motor)
	})
	ui.Bind("coolantHeat", func(spindlePower string,roughTime string,finishTime string,chipremoval string,power1 string,num1 string,power2 string,num2 string,power3 string,num3 string,power4 string,num4 string,surface string,temperature string) {
		worker.RoughTime ,_=strconv.ParseFloat(roughTime, 64)
		worker.FinishTime ,_=strconv.ParseFloat(finishTime, 64)
		worker.Chipremoval ,_=strconv.ParseFloat(chipremoval, 64)
		worker.SpindlePower,_=strconv.ParseFloat(spindlePower, 64)

		pump.Power1,_=strconv.ParseFloat(power1, 64)
		pump.Power2,_=strconv.ParseFloat(power2, 64)
		pump.Power3,_=strconv.ParseFloat(power3, 64)
		pump.Power4,_=strconv.ParseFloat(power4, 64)

		pump.Num1,_=strconv.ParseFloat(num1, 64)
		pump.Num2,_=strconv.ParseFloat(num2, 64)
		pump.Num3,_=strconv.ParseFloat(num3, 64)
		pump.Num4,_=strconv.ParseFloat(num4, 64)

		coolant.Surface ,_=strconv.ParseFloat(surface, 64)
		coolant.Temperature,_=strconv.ParseFloat(temperature, 64)
		
		if worker.FinishTime+worker.RoughTime>1 {
			ui.Eval(`alert("加工總時間不該超過1");`)
		}
		ui.Eval(`document.getElementById("workerW").innerText = `+strconv.FormatFloat(worker.q(), 'E', 6, 64))		
		ui.Eval(`document.getElementById("pumpW").innerText = `+strconv.FormatFloat(pump.q(), 'E', 6, 64))		
		ui.Eval(`document.getElementById("coolantW").innerText = `+strconv.FormatFloat(coolant.q(), 'E', 6, 64))		
		ui.Eval(`document.getElementById("workW").innerText = `+strconv.FormatFloat(worker.q()+pump.q()+coolant.q(), 'E', 6, 64))		
		
	})
	ui.Bind("movingHeat", func(weight string,axiaload string,v string,bsLead string,bsEff string,gwR string,gwU string,gwType string) {
		moving.Weight ,_=strconv.ParseFloat(weight, 64)
		moving.Axiaload ,_=strconv.ParseFloat(axiaload, 64)
		moving.V ,_=strconv.ParseFloat(v, 64)
		moving.Tballscrew.Lead ,_=strconv.ParseFloat(bsLead, 64)
		moving.Tballscrew.Effectiveness ,_=strconv.ParseFloat(bsEff, 64)
		moving.Tguide.R ,_=strconv.ParseFloat(gwR, 64)
		moving.Tguide.U ,_=strconv.ParseFloat(gwU, 64)
		moving.Tguide.Type ,_=strconv.ParseFloat(gwType, 64)
		
		ui.Eval(`document.getElementById("bsW").innerText = `+strconv.FormatFloat(moving.screwHeat(), 'E', 6, 64))		
		ui.Eval(`document.getElementById("gwW").innerText = `+strconv.FormatFloat(moving.guideHeat(), 'E', 6, 64))	
		ui.Eval(`document.getElementById("bsRpm").innerText = `+strconv.FormatFloat(moving.Tballscrew.rpm, 'E', 6, 64))		
		ui.Eval(`document.getElementById("bsTorque").innerText = `+strconv.FormatFloat(moving.Tballscrew.torque, 'E', 6, 64))		
		fmt.Println(moving)
	})

	<-ui.Done()
}

//in call
func read(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}