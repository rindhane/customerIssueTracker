const signInForm = document.getElementById('signInForm');
signInForm.addEventListener('submit',validateCredentials)


async function validateCredentials(event)
{
    event.preventDefault();
    const JSONResult = await formDatatoJson(event.target);
    const {status:stat, data:responseTxt}= await postJsonData(url="/checkAuth",JsonData=JSONResult);
    if(stat==200){
        const tokenData = JSON.parse(responseTxt);
        storeInLocalStorage("qdasAuthToken",tokenData.tokenString);
        await setCookie("authToken",tokenData.tokenString);
        onSuccessGoToDefaultPage();
    }
}

async function onSuccessGoToDefaultPage(){
    window.location.href = "/UserDashboard";
}


