async function onloadModal(){
    let signUpfieldActive=false;
    const otpDisplayWrapper = document.getElementById("otpDisplayWrapper");
    const emailFieldElement = document.getElementById("emailSignUp");
    const signUpSubmission = document.getElementById("signUpSubmission");
    const otpFieldElement = document.getElementById("otpSignUp");
    emailFieldElement.addEventListener("input", (event)=>
        {
        actionOnInput(event,otpDisplayWrapper, className="hideDisplayContainer");
        }
        );
    otpFieldElement.addEventListener("input", (event)=>
    {
    actionOnInput(event,signUpSubmission,className="hideDisplayContainer");
    }
    );
}


async function actionOnInput(event,elem,className){
    const inputValue = event.target.value;
    if (inputValue==""){
        elem.classList.add(className);
        return 
    }
    elem.classList.remove(className);
}

async function submitAccountDetails(){
    const otpFieldElement = document.getElementById("otpSignUp");
    const emailFieldElement = document.getElementById("emailSignUp");
    const jsonPayload = {
        email:emailFieldElement.value,
        password:otpFieldElement.value
    }
    const response = await postJsonData("/OtpAuth",jsonPayload);
    const result = JSON.parse(response.data);
    confirm(result.remark);
}
