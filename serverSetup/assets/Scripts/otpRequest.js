
async function generateOTPRequest() {
    const emailFieldElement = document.getElementById("emailSignUp");
    const email=emailFieldElement.value;
    const jsonPayload = {
        email:email,
    }
    const response = await postJsonData("/generateOTP",jsonPayload);
    const result = JSON.parse(response.data);
    confirm(result.remark);
}