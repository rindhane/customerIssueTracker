function applyAttributes(elem,classArray){
    let classString = "";
    for (let i=0; i <classArray.length; i++){
        classString=classString+" "+ classArray[i]+" ";
    }
    elem.setAttribute('class',classString);
    return elem;
}

function generateDomElement(typ, classDetail, innerText){
    const elem = document.createElement(typ);
    applyAttributes(elem,classDetail);
    elem.innerText=innerText;
    return elem;
}


//preparing payloads 
async function formDatatoJson(form){
    data = new FormData(form);
    const result= {};
    result.email= data.get("email");
    result.password =  await digestMessage(data.get("password"));
    return result;
}


//sending post request to server
async function postJsonData(url , Jsondata, authKey="qdasAuthToken"){ 
    let options={
      method: 'POST',
      headers: {
        Accept: 'application/json',
        "Content-Type": 'application/json',
        authToken:getKeyValueFromStorage(authKey)
      },
      body: JSON.stringify(Jsondata),
      cache: 'default',
      redirect:'manual',
    }
    let response = await fetch(url, options);
    const status = response.status  
    const data = await response.text();
    return { status:status, data:data }
  };

function responseToJson(textData){
    return JSON.parse(textData);
}



//local storage functions

function storeInLocalStorage(key,value){
    localStorage.setItem(key,value);
    return true;
}

function getKeyValueFromStorage(key){
    const result = localStorage.getItem(key);
    return result;
}

// local cookie functions

async function setCookie(name,value){
    document.cookie=`${name}=${value}`
    return true
}


// field Validators
async function validateField(inputElement){
    if (inputElement.value=""){
        return false;
    }
}

//crypto functions 


async function digestMessage(message){
    const encoder = new TextEncoder();
    const msgUint8 = encoder.encode(message); //encoded as utf-8
    const hashBuffer = await crypto.subtle.digest("SHA-256",msgUint8);
    const hashArray=Array.from(new Uint8Array(hashBuffer)); //get byte array
    const hashHex= hashArray.map((b)=>b.toString(16).padStart(2,"0")).join("");
    return hashHex;
    //should modify to return in base64String : 
    //ref:https://developer.mozilla.org/en-US/docs/Web/API/btoa
    //https://stackoverflow.com/questions/9267899/arraybuffer-to-base64-encoded-string
} 

