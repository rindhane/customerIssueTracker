const IssueTable= document.getElementById("IssueTable");
const modalContainer = document.getElementById("modalBackground");
//const closeModalButton = document.getElementById("close");
const sideBar = document.getElementById("sideBar")
const templateHTML = document.getElementById("modal_template");

async function getIssuesUser(url="/UserIssues"){
    const {status:stat, data:rawText}=await postJsonData(url,"give userData");
    const issuesBlock = responseToJson(rawText);
    return issuesBlock;
}

function generateDOMattributes(txt){
    const classMap= new Map() ;
    classMap.set("shortDesc",["item", "descBlock"]);
    classMap.set("lastAction",["item"]);
    classMap.set("status",["item"]);
    classMap.set("sn",["item", "shortBlock"]);
    classMap.set("itemSection" , ["itemSection",])
    const DomMap = new Map();
    DomMap.set("shortDesc","div");
    DomMap.set("lastAction","div");
    DomMap.set("status","div");
    DomMap.set("sn","div");
    DomMap.set("itemSection","div");
    return {
        classArray:classMap.get(txt),
        domTyp:DomMap.get(txt),
    };
}

async function renderIssueTable(tableElem, issueBlock){
    for(let i=0; i< issueBlock.length; i++){
        tableElem.appendChild(generateItemSectionDOM(issueBlock[i],i+1));
    }
    return tableElem;
}

function generateItemSectionDOM(issueObj, index){
    const keyArray=Object.keys(issueObj).filter(
        (key)=>{
            if(key==="issueID"){
                return false;
            }
            return true;
        }
    )
    const parentAttrib=generateDOMattributes('itemSection');
    const sectionElement=generateDomElement(parentAttrib.domTyp,parentAttrib.classArray,"");
    //adding 'sn' dom element 
    let attrib = generateDOMattributes('sn');
    let chilElem=generateDomElement(attrib.domTyp,attrib.classArray,index);
    sectionElement.appendChild(chilElem);
    keyArray.forEach( // loop to generate childElements and attach to sectionElement;
        element => {
        let attrib = generateDOMattributes(element);
        let chilElem=generateDomElement(attrib.domTyp,attrib.classArray,issueObj[element]);
        sectionElement.appendChild(chilElem);
    }); 
    return sectionElement;
}

async function raiseModal(){
    modalContainer.innerHTML = templateHTML.innerHTML.slice();
    modalContainer.style.display="flex";
    try {
        onloadModal();
    }catch{

    }
};
function closeModal(){
    modalContainer.style.display="none";
};

async function saveNewIssueDetails(){
    const issueType= document.getElementById("issueType");
    const issueDescription = document.getElementById("issueDescription");
    const issueInput ={
        type:parseInt(issueType.value,10),
        description:issueDescription.value,
    }
    const result= await postJsonData("/newIssueRaise",issueInput,);
    if (result.status ==200){
        closeModal();
        confirm("issue submitted");
    }else {
        confirm('something went wrong');
    }
}

async function renderSidebar(){
    const tokenString = getKeyValueFromStorage("qdasAuthToken");
    const splitArray=tokenString.split('.'); 
    const claimString =splitArray[1]; // claims block
    const decoded = atob(claimString);
    const claimsObj=JSON.parse(decoded);
    console.log(claimsObj.level);
}




