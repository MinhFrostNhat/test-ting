"use client";

import React, { useCallback, useEffect, useRef, useState } from 'react';
import axios from "axios";
import './page.css';
import { useImmer } from "use-immer";
import DataRow from "../components/DataRow";

export interface Kol {
	kolID: number;
	userProfileID: number;
	language: string;
	education: string;
	expectedSalary: number;
	expectedSalaryEnable: boolean;
	channelSettingTypeID: number;
	idFrontURL: string;
	idBackURL: string;
	portraitURL: string;
	rewardID: number;
	paymentMethodID: number;
	testimonialsID: number;
	verificationStatus: boolean;
	enabled: boolean;
	activeDate: Date;
	active: boolean;
	createdBy: string;
	createdDate: Date;
	modifiedBy: string;
	modifiedDate: Date;
	isRemove: boolean;
	isOnBoarding: boolean;
	code: string;
	portraitRightURL: string;
	portraitLeftURL: string;
	livenessStatus: boolean;
}

// Create an Axios instance with a base URL
const api = axios.create({
	baseURL: "http://localhost:8081/api/v1/"
});

const Page = () => {
	const [kols, setKols] = useState<Kol[]>([]);
	const [pagination, updatePagination] = useImmer({
		pageIndex: 1,
		pageSize: 15
	});
	const kolRef = useRef<HTMLDivElement>(null);

	useEffect(() => {
		const fetchKols = async () => {
			try {
				const res = await api.get<{ kol: Kol[] }>("kols", {
					params: pagination
				});
				setKols(res.data.kol);
			} catch (error) {
				console.error("Error fetching KOL data:", error);
			}
		};

		fetchKols();
	}, [pagination]);

	const scrollKols = useCallback((scrollOffset: number) => {
		kolRef.current?.scrollBy({
			top: scrollOffset,
			behavior: 'smooth' // Optional: adds smooth scrolling
		});
	}, []);

	return (
		<>
			<h1 className='header'>Implement component </h1>
			<div style={{ position: 'relative', maxHeight: '100vh', overflow: "hidden" }}>
				<div style={{ width: '100%', height: 50, position: "absolute", display: "flex", justifyContent: "space-between" }}>
					<button onClick={() => scrollKols(-1512)} className="scroll-button left">
						Scroll Up
					</button>
					<button onClick={() => scrollKols(1512)} className="scroll-button right">
						Scroll Down
					</button>
				</div>
				<div style={{ maxHeight: '100vh', overflow: "scroll" }} ref={kolRef} id="kolsContainer">
					{kols.map(kol => (
						<DataRow key={kol.kolID} kol={kol} />
					))}
				</div>
			</div>
		</>
	);
};


export default Page;