// Testtt --
func (webService *WebService) Testtt(c *gin.Context) {
	fileNameList := make(map[int]string)
	for i := 0; i < 15; i++ {
		fileNameList[i] = fmt.Sprintf("photo/%v.jpg", strconv.Itoa(i))
	}
	for _, k := range fileNameList {
		fmt.Println(k)
	}
	file, err := os.Open("photo/0.jpg")
	if err != nil {
		fmt.Println("err = ", err.Error())
		return
	}

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("err = ", err.Error())
		return
	}
	defer file.Close()
	b1 := make([]byte, fi.Size())
	file.Read(b1)
}


// Test22 - test method
func (webService *WebService) Test22(c *gin.Context) {
	var compareNotify f2fmodel.CompareNotify2

	iinList, err := readLines()
	if err != nil {
		fmt.Println("error = ", "-111, ", err.Error())
		return
	}

	resCnt := 0
	for i, clientIIN := range iinList {
		if i >= 41 {
			return
		}
		ocrmURL := fmt.Sprintf(webService.Config.OCRMWS, clientIIN)
		req, err := http.NewRequest("GET", ocrmURL, nil)
		if err != nil {
			fmt.Println("error = ", "-222, ", err.Error())
			break
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("error = ", "-333, ", err.Error())
			break
		}

		htmlData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error = ", "-444, ", err.Error())
			break
		}
		if len(htmlData) < 60 {
			fmt.Println("error = ", "-555, OCRM: ", string(htmlData), "iin = ", clientIIN)
			break
		}
		defer resp.Body.Close()

		compareNotify.IIN = clientIIN
		compareNotify.PhotoB64 = b64.StdEncoding.EncodeToString(htmlData)
		vendorsList := make(map[string]strfmt.UUID4)
		vendorsList["1"] = "d9eaebc5-c473-4f0e-beea-a008db795f87"
		vendorsList["2"] = "a13da182-9825-49bf-8959-edb4aacc97ce"
		vendorsList["3"] = "f79e71d9-0f16-4e4d-9de9-8f40321b0d9e"
		vendorsList["4"] = "8171a931-e8b2-4d9a-81b5-f140f24ce6ea"
		//vendorsList["5"] = "463f0cd5-7e7e-43c4-9526-c5274e9732da"
		//vendorsList["6"] = "5d5bc3e1-b905-478f-8cc7-f4d35138c230"

		fmt.Println(compareNotify.IIN)
		for _, vList := range vendorsList {
			compareNotify.VendorID = string(vList)
			go webService.Kafka.Send(compareNotify, webService.Config.CompareTopic)
			resCnt = resCnt + 1
		}

		/*
			var vendorsList [6]string
			vendorsList[0] = "d9eaebc5-c473-4f0e-beea-a008db795f87"
			vendorsList[1] = "a13da182-9825-49bf-8959-edb4aacc97ce"
			vendorsList[2] = "f79e71d9-0f16-4e4d-9de9-8f40321b0d9e"
			vendorsList[3] = "8171a931-e8b2-4d9a-81b5-f140f24ce6ea"
			vendorsList[4] = "463f0cd5-7e7e-43c4-9526-c5274e9732da"
			vendorsList[5] = "5d5bc3e1-b905-478f-8cc7-f4d35138c230"
			/*vendorsList, err := webService.ManagerDB.GetVendors()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "-11, " + err.Error()})
				webService.Logger.Error("VerifyResult", "FaceRecognition.Core", -11, "", err.Error(), "", "", nil)
				return
			}*/
		/*for _, vendor := range vendorsList {
			compareNotify.VendorID = vendor
			go webService.Kafka.Send(compareNotify, webService.Config.CompareTopic)
		}*/

		/*var vendorsCompare dbmodel.VendorsCompare
			vendorsCompare.VendorID = "d9eaebc5-c473-4f0e-beea-a008db795f87"
			vendorsCompare.Idempotency = ""
			vendorsCompare.Similarity = 0
			vendorsCompare.CompareResultID = "d9eaebc5-c473-4f0e-beea-a008db795f87"
			err = webService.Manager.CreateVendorsCompare(&vendorsCompare)
			if err != nil {
				fmt.Println("err=", err.Error())
			}

		}
		c.JSON(200, compareNotify)*/
	}

	fmt.Println("count = ", resCnt)
}


func (webService *WebService) Test23(c *gin.Context) {
	fileNameList := make(map[int]string)
	fileBodyList := make(map[string][]byte)
	for i := 0; i < 4; i++ {
		fileNameList[i] = fmt.Sprintf("photo/%v.jpg", strconv.Itoa(i))
	}

	for _, fileName := range fileNameList {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("err = ", err.Error())
			return
		}

		fi, err := file.Stat()
		if err != nil {
			fmt.Println("err = ", err.Error())
			return
		}
		defer file.Close()
		htmlData := make([]byte, fi.Size())
		file.Read(htmlData)
		fileBodyList[fileName] = htmlData
		fmt.Println("filePath = ", fileName)
	}
	var compareNotify f2fmodel.CompareNotify2
	var clIin dbmodel.ClientIin
	/*htmlData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("error = ", "-111, ", err.Error())
		c.JSON(400, gin.H{"error": "-111, " + err.Error()})
	}
	defer c.Request.Body.Close()
	compareNotify.PhotoB64 = b64.StdEncoding.EncodeToString(htmlData)*/

	tx := webService.Manager.DB.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "-1, " + tx.Error.Error()})
		webService.Logger.Error("VerifyResultNotify", "FaceRecognition.Core", -1, "", tx.Error.Error(), "", "", nil)
		return
	}
	iinList, err := webService.Manager.GetIINList(&clIin, tx)
	if err != nil {
		fmt.Println("error = ", "-111, ", err.Error())
		return
	}

	err = tx.Commit().Error
	if err != nil {
		webService.Logger.Error("FaceRecogniotion.Core", "CompareResult", -12, err.Error(), "Не корректный запрос", "", "", nil)
		c.JSON(http.StatusBadRequest, gin.H{"error": "-12, " + err.Error()})
		return
	}

	resCnt := 0
	for i, iinInfo := range iinList {
		//resCnt = resCnt + 1
		if i >= 3 {
			fmt.Println("i ============== ", i)
			fmt.Println("count = ", resCnt)
			return
		}
		clientIIN := iinInfo.IIN

		compareNotify.IIN = clientIIN
		vendorsList := make(map[string]strfmt.UUID4)
		vendorsList["1"] = "d9eaebc5-c473-4f0e-beea-a008db795f87"
		vendorsList["2"] = "a13da182-9825-49bf-8959-edb4aacc97ce"
		//vendorsList["3"] = "f79e71d9-0f16-4e4d-9de9-8f40321b0d9e"
		vendorsList["4"] = "8171a931-e8b2-4d9a-81b5-f140f24ce6ea"

		for _, fileBody := range fileBodyList {
			//resCnt = resCnt + 1

			compareNotify.PhotoB64 = b64.StdEncoding.EncodeToString(fileBody)

			for _, vList := range vendorsList {
				resCnt = resCnt + 1
				compareNotify.VendorID = string(vList)
				webService.Kafka.Send(compareNotify, webService.Config.CompareTopic)
				//fmt.Println("iiii = ", i, "iin = ", compareNotify.IIN)
			}
		}

	}

}