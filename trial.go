package main

import (
    "fmt"
    "gltf-master"
    "flag"
)

func main() {
    var doc *gltf.Document
    var doc1 *gltf.Document
    
    var name string
    flag.StringVar(&name,"opt","","Usage")
    flag.Parse()
    fmt.Println(name)
    //For Capital*.gltf
    capitalSlice := []byte{67, 97, 112, 105, 116, 97, 108, 47, 46, 103, 108, 116, 102} 
    
    //For Small*.gltf
    smallSlice := []byte{83, 109, 97, 108, 108, 47, 46, 103, 108, 116, 102} 
    
    var modelPool [100]*gltf.Document 
    
    var mystring1 string          	

    //var inputString string
    var prevIsSpace int
    for index, s := range name { 
        if s==32 || s==95{
            prevIsSpace=1
            continue
        }
        if index==0 ||prevIsSpace==1{
                capitalSlice[7]=byte(s)
                mystring1 = string(capitalSlice)
                prevIsSpace=0                
            }else{
                smallSlice[5]=byte(s)
                mystring1 = string(smallSlice)
            }
        modelPool[index],_=gltf.Open(mystring1)
    } 
    
    doc = modelPool[0]

    printEnable:=0
    
    if(printEnable==1){
        doc.Meshes[0].Name="Cube"
        fmt.Println(doc.Meshes[0].Name)
        fmt.Println(doc.Meshes[0])
        fmt.Println(len(doc1.Scenes[0].Nodes))
    }
    
    var meshTemp *uint32
    meshTemp = nil 
    
        //Mesh Primitive Max Calc 
        meshPrimitiveMax := uint32(0)
        for kOut := range doc.Meshes{
            for loopIter := range doc.Meshes[kOut].Primitives[0].Attributes {
                //fmt.Println("inInnerLoop ", meshPrimitiveMax, "-", kOut,"-", doc.Meshes[0].Primitives[0].Attributes[loopIter])
                if(meshPrimitiveMax<doc.Meshes[kOut].Primitives[0].Attributes[loopIter]) {
                    meshPrimitiveMax=doc.Meshes[kOut].Primitives[0].Attributes[loopIter]
                }
            }
      
            if(doc.Meshes[kOut].Primitives[0].Indices != meshTemp) {
                if(meshPrimitiveMax<*doc.Meshes[kOut].Primitives[0].Indices) {
                    meshPrimitiveMax=*doc.Meshes[kOut].Primitives[0].Indices
                }
            }
        }

        //Place to adjust the position of the characters
        doc.Nodes[0].Translation[0]=0
        doc.Nodes[0].Translation[1]=0
        doc.Nodes[0].Translation[2]=0
        //copier.Copy(&doc1, &doc2) //Deep Copy
    for loopIter:=0; loopIter<len(name)-1; loopIter++ {      
        // This is the place where position of the Alphabets is needed to be implemented   
        doc1 = modelPool[loopIter+1]    
        if doc1 == nil{
            continue
        }
        //Place to adjust the position of the characters
        doc1.Nodes[0].Translation[0]=float64(loopIter+1)
        doc1.Nodes[0].Translation[1]=0
        doc1.Nodes[0].Translation[2]=0
        
        //Start of actucal Merge
        for n := 0; n <len(doc1.Scenes[0].Nodes); n++ {
            
            //Scenes.Nodes related Code Changes
            doc.Scenes[0].Nodes = append(doc.Scenes[0].Nodes, uint32(len(doc.Scenes[0].Nodes)))
        
            //Nodes related Changes 
            doc.Nodes = append(doc.Nodes, doc1.Nodes[n])
            *doc.Nodes[len(doc.Nodes)-1].Mesh=uint32(len(doc.Nodes)-1)  
        }
        
        //Buffers related code 
        doc.Buffers = append(doc.Buffers, doc1.Buffers[0])
    
        //BufferViews Changes
        for n:=0 ; n<len(doc1.BufferViews); n++ {   
            //doc1.BufferViews[n].Buffer=uint32(len(doc.Buffers)-1)
            doc.BufferViews = append(doc.BufferViews, doc1.BufferViews[n])
            doc.BufferViews[len(doc.BufferViews)-1].Buffer=uint32(len(doc.Buffers)-1)
        }
    
        //Accessors related changes
        a:=uint32(len(doc.Accessors))
        for n:=0 ; n<len(doc1.Accessors); n++ {
            //*doc1.Accessors[n].BufferView+=a
            doc.Accessors = append(doc.Accessors, doc1.Accessors[n])
            *doc.Accessors[len(doc.Accessors)-1].BufferView+=a
        }
        
        //Mesh related part two
        meshPrimitiveMaxTemp:=meshPrimitiveMax
        for n:=0 ; n<len(doc1.Meshes); n++ {
            //Appendix 1
            doc.Meshes = append(doc.Meshes, doc1.Meshes[n])
            
            for kIn := range doc.Meshes[len(doc.Meshes)-1].Primitives[0].Attributes {
               if printEnable==1{
                    fmt.Println("meshPrimitiveMax updated 1: ",meshPrimitiveMax,"-",doc1.Meshes[n].Primitives[0].Attributes[kIn],"-",meshPrimitiveMaxTemp)    
               }
               doc.Meshes[len(doc.Meshes)-1].Primitives[0].Attributes[kIn]+=(meshPrimitiveMax+1)
                   if(meshPrimitiveMaxTemp<doc.Meshes[len(doc.Meshes)-1].Primitives[0].Attributes[kIn]){
                    meshPrimitiveMaxTemp=doc.Meshes[len(doc.Meshes)-1].Primitives[0].Attributes[kIn]                
                } 
                if printEnable==1{
                    fmt.Println("meshPrimitiveMax updated 2: ",meshPrimitiveMax,"-",doc1.Meshes[n].Primitives[0].Attributes[kIn],"-",meshPrimitiveMaxTemp)    
                }
            }
        
            if(doc.Meshes[len(doc.Meshes)-1].Primitives[0].Indices != meshTemp) {
                *doc.Meshes[len(doc.Meshes)-1].Primitives[0].Indices+=(meshPrimitiveMax+1)
                if(meshPrimitiveMaxTemp<*doc.Meshes[len(doc.Meshes)-1].Primitives[0].Indices){
                    meshPrimitiveMaxTemp=*doc.Meshes[len(doc.Meshes)-1].Primitives[0].Indices
                }    
                if printEnable==1{
                    fmt.Println("meshPrimitiveMax updated 3: ",meshPrimitiveMax,"-",*doc.Meshes[len(doc.Meshes)-1].Primitives[0].Indices,"-",meshPrimitiveMaxTemp) 
                }
            }
            
            meshPrimitiveMax=meshPrimitiveMaxTemp
        }
    	//End of actual Merge
        
        if printEnable==1 {
            fmt.Println("Attributes : ", doc.Meshes[0].Primitives[0].Attributes)
            fmt.Println("normal : ", doc.Meshes[0].Primitives[0].Attributes["NORMAL"])
            fmt.Println("Length : ", len(doc.Meshes[0].Primitives[0].Attributes))  
        }          
    }
    
    if err := gltf.Save(doc, "LAVANYATRIPATI.gltf"); err != nil {
        panic(err)
    }
}