import React from 'react';
import { Kol } from "@/app/page";  // Ensure this import correctly points to the Kol type definition

interface DataRowProps {
        kol: Kol;
}

const DataRow: React.FC<DataRowProps> = ({ kol }) => {
    return (
        <div className="container">
            <h2>KOL Data</h2>
            <div className="data-row"><span>KOL ID:</span> {kol.kolID}</div>
            <div className="data-row"><span>User Profile ID:</span> {kol.userProfileID}</div>
            <div className="data-row"><span>Language:</span> {kol.language}</div>
            <div className="data-row"><span>Education:</span> {kol.education}</div>
            <div className="data-row"><span>Expected Salary:</span> ${kol.expectedSalary.toLocaleString()}</div>
            <div className="data-row"><span>Expected Salary Enabled:</span> {kol.expectedSalaryEnable.toString()}</div>
            <div className="data-row"><span>Channel Setting Type ID:</span> {kol.channelSettingTypeID}</div>
            <div className="data-row"><span>ID Front URL:</span> <img src={kol.idFrontURL} alt="ID Front"/></div>
            <div className="data-row"><span>ID Back URL:</span> <img src={kol.idBackURL} alt="ID Back"/></div>
            <div className="data-row"><span>Portrait URL:</span> <img src={kol.portraitURL} alt="Portrait"/></div>
            <div className="data-row"><span>Reward ID:</span> {kol.rewardID}</div>
            <div className="data-row"><span>Payment Method ID:</span> {kol.paymentMethodID}</div>
            <div className="data-row"><span>Testimonials ID:</span> {kol.testimonialsID}</div>
            <div className="data-row"><span>Verification Status:</span> {kol.verificationStatus.toString()}</div>
            <div className="data-row"><span>Enabled:</span> {kol.enabled.toString()}</div>
            <div className="data-row"><span>Active Date:</span> {kol.activeDate.toString()}</div>
            <div className="data-row"><span>Active:</span> {kol.active.toString()}</div>
            <div className="data-row"><span>Created By:</span> {kol.createdBy}</div>
            <div className="data-row"><span>Created Date:</span> {kol.createdDate.toString()}</div>
            <div className="data-row"><span>Modified By:</span> {kol.modifiedBy}</div>
            <div className="data-row"><span>Modified Date:</span> {kol.modifiedDate.toString()}</div>
            <div className="data-row"><span>Is Remove:</span> {kol.isRemove.toString()}</div>
            <div className="data-row"><span>Is OnBoarding:</span> {kol.isOnBoarding.toString()}</div>
            <div className="data-row"><span>Code:</span> {kol.code}</div>
            <div className="data-row"><span>Portrait Right URL:</span> <img src={kol.portraitRightURL}
                                                                            alt="Portrait Right Image"/></div>
            <div className="data-row"><span>Portrait Left URL:</span> <img src={kol.portraitLeftURL}
                                                                           alt="Portrait Left Image"/></div>
            <div className="data-row"><span>Liveness Status:</span> {kol.livenessStatus.toString()}</div>
        </div>

    );
};

export default DataRow;
