package ctpstruct

type TThostFtdcTraderIDType [21]byte
type TThostFtdcInvestorIDType [13]byte
type TThostFtdcBrokerIDType [11]byte
type TThostFtdcBrokerAbbrType [9]byte
type TThostFtdcBrokerNameType [81]byte
type TThostFtdcExchangeInstIDType [31]byte
type TThostFtdcOrderRefType [13]byte
type TThostFtdcParticipantIDType [11]byte
type TThostFtdcUserIDType [16]byte
type TThostFtdcPasswordType [41]byte
type TThostFtdcClientIDType [11]byte
type TThostFtdcInstrumentIDType [31]byte
type TThostFtdcInstrumentCodeType [31]byte
type TThostFtdcMarketIDType [31]byte
type TThostFtdcProductNameType [21]byte
type TThostFtdcExchangeIDType [9]byte
type TThostFtdcExchangeNameType [61]byte
type TThostFtdcExchangeAbbrType [9]byte
type TThostFtdcExchangeFlagType [2]byte
type TThostFtdcMacAddressType [21]byte
type TThostFtdcSystemIDType [21]byte
type TThostFtdcExchangePropertyType byte
type TThostFtdcDateType [9]byte
type TThostFtdcTimeType [9]byte
type TThostFtdcLongTimeType [13]byte
type TThostFtdcInstrumentNameType [21]byte
type TThostFtdcSettlementGroupIDType [9]byte
type TThostFtdcOrderSysIDType [21]byte
type TThostFtdcTradeIDType [21]byte
type TThostFtdcCommandTypeType [65]byte
type TThostFtdcIPAddressType [16]byte
type TThostFtdcIPPortType int64
type TThostFtdcProductInfoType [11]byte
type TThostFtdcProtocolInfoType [11]byte
type TThostFtdcBusinessUnitType [21]byte
type TThostFtdcDepositSeqNoType [15]byte
type TThostFtdcIdentifiedCardNoType [51]byte
type TThostFtdcIdCardTypeType byte
type TThostFtdcOrderLocalIDType [13]byte
type TThostFtdcUserNameType [81]byte
type TThostFtdcPartyNameType [81]byte
type TThostFtdcErrorMsgType [81]byte
type TThostFtdcFieldNameType [2049]byte
type TThostFtdcFieldContentType [2049]byte
type TThostFtdcSystemNameType [41]byte
type TThostFtdcContentType [501]byte
type TThostFtdcInvestorRangeType byte
type TThostFtdcDepartmentRangeType byte
type TThostFtdcDataSyncStatusType byte
type TThostFtdcBrokerDataSyncStatusType byte
type TThostFtdcExchangeConnectStatusType byte
type TThostFtdcTraderConnectStatusType byte
type TThostFtdcFunctionCodeType byte
type TThostFtdcBrokerFunctionCodeType byte
type TThostFtdcOrderActionStatusType byte
type TThostFtdcOrderStatusType byte
type TThostFtdcOrderSubmitStatusType byte
type TThostFtdcPositionDateType byte
type TThostFtdcPositionDateTypeType byte
type TThostFtdcTradingRoleType byte
type TThostFtdcProductClassType byte
type TThostFtdcInstLifePhaseType byte
type TThostFtdcDirectionType byte
type TThostFtdcPositionTypeType byte
type TThostFtdcPosiDirectionType byte
type TThostFtdcSysSettlementStatusType byte
type TThostFtdcRatioAttrType byte
type TThostFtdcHedgeFlagType byte
type TThostFtdcBillHedgeFlagType byte
type TThostFtdcClientIDTypeType byte
type TThostFtdcOrderPriceTypeType byte
type TThostFtdcOffsetFlagType byte
type TThostFtdcForceCloseReasonType byte
type TThostFtdcOrderTypeType byte
type TThostFtdcTimeConditionType byte
type TThostFtdcVolumeConditionType byte
type TThostFtdcContingentConditionType byte
type TThostFtdcActionFlagType byte
type TThostFtdcTradingRightType byte
type TThostFtdcOrderSourceType byte
type TThostFtdcTradeTypeType byte
type TThostFtdcPriceSourceType byte
type TThostFtdcInstrumentStatusType byte
type TThostFtdcInstStatusEnterReasonType byte
type TThostFtdcOrderActionRefType int64
type TThostFtdcInstallCountType int64
type TThostFtdcInstallIDType int64
type TThostFtdcErrorIDType int64
type TThostFtdcSettlementIDType int64
type TThostFtdcVolumeType int64
type TThostFtdcFrontIDType int64
type TThostFtdcSessionIDType int64
type TThostFtdcSequenceNoType int64
type TThostFtdcCommandNoType int64
type TThostFtdcMillisecType int64
type TThostFtdcVolumeMultipleType int64
type TThostFtdcTradingSegmentSNType int64
type TThostFtdcRequestIDType int64
type TThostFtdcYearType int64
type TThostFtdcMonthType int64
type TThostFtdcBoolType int64
type TThostFtdcPriceType float64
type TThostFtdcCombOffsetFlagType [5]byte
type TThostFtdcCombHedgeFlagType [5]byte
type TThostFtdcRatioType float64
type TThostFtdcMoneyType float64
type TThostFtdcLargeVolumeType float64
type TThostFtdcSequenceSeriesType int16
type TThostFtdcCommPhaseNoType int16
type TThostFtdcSequenceLabelType [2]byte
type TThostFtdcUnderlyingMultipleType float64
type TThostFtdcPriorityType int64
type TThostFtdcContractCodeType [41]byte
type TThostFtdcCityType [51]byte
type TThostFtdcIsStockType [11]byte
type TThostFtdcChannelType [51]byte
type TThostFtdcAddressType [101]byte
type TThostFtdcZipCodeType [7]byte
type TThostFtdcTelephoneType [41]byte
type TThostFtdcFaxType [41]byte
type TThostFtdcMobileType [41]byte
type TThostFtdcEMailType [41]byte
type TThostFtdcMemoType [161]byte
type TThostFtdcCompanyCodeType [51]byte
type TThostFtdcWebsiteType [51]byte
type TThostFtdcTaxNoType [31]byte
type TThostFtdcBatchStatusType byte
type TThostFtdcPropertyIDType [33]byte
type TThostFtdcPropertyNameType [65]byte
type TThostFtdcLicenseNoType [51]byte
type TThostFtdcAgentIDType [13]byte
type TThostFtdcAgentNameType [41]byte
type TThostFtdcAgentGroupIDType [13]byte
type TThostFtdcAgentGroupNameType [41]byte
type TThostFtdcReturnStyleType byte
type TThostFtdcReturnPatternType byte
type TThostFtdcReturnLevelType byte
type TThostFtdcReturnStandardType byte
type TThostFtdcMortgageTypeType byte
type TThostFtdcInvestorSettlementParamIDType byte
type TThostFtdcExchangeSettlementParamIDType byte
type TThostFtdcSystemParamIDType byte
type TThostFtdcTradeParamIDType byte
type TThostFtdcSettlementParamValueType [256]byte
type TThostFtdcCounterIDType [33]byte
type TThostFtdcInvestorGroupNameType [41]byte
type TThostFtdcBrandCodeType [257]byte
type TThostFtdcWarehouseType [257]byte
type TThostFtdcProductDateType [41]byte
type TThostFtdcGradeType [41]byte
type TThostFtdcClassifyType [41]byte
type TThostFtdcPositionType [41]byte
type TThostFtdcYieldlyType [41]byte
type TThostFtdcWeightType [41]byte
type TThostFtdcSubEntryFundNoType int64
type TThostFtdcFileIDType byte
type TThostFtdcFileNameType [257]byte
type TThostFtdcFileTypeType byte
type TThostFtdcFileFormatType byte
type TThostFtdcFileUploadStatusType byte
type TThostFtdcTransferDirectionType byte
type TThostFtdcUploadModeType [21]byte
type TThostFtdcAccountIDType [13]byte
type TThostFtdcBankFlagType [4]byte
type TThostFtdcBankAccountType [41]byte
type TThostFtdcOpenNameType [61]byte
type TThostFtdcOpenBankType [101]byte
type TThostFtdcBankNameType [101]byte
type TThostFtdcPublishPathType [257]byte
type TThostFtdcOperatorIDType [65]byte
type TThostFtdcMonthCountType int64
type TThostFtdcAdvanceMonthArrayType [13]byte
type TThostFtdcDateExprType [1025]byte
type TThostFtdcInstrumentIDExprType [41]byte
type TThostFtdcInstrumentNameExprType [41]byte
type TThostFtdcSpecialCreateRuleType byte
type TThostFtdcBasisPriceTypeType byte
type TThostFtdcProductLifePhaseType byte
type TThostFtdcDeliveryModeType byte
type TThostFtdcLogLevelType [33]byte
type TThostFtdcProcessNameType [257]byte
type TThostFtdcOperationMemoType [1025]byte
type TThostFtdcFundIOTypeType byte
type TThostFtdcFundTypeType byte
type TThostFtdcFundDirectionType byte
type TThostFtdcFundStatusType byte
type TThostFtdcBillNoType [15]byte
type TThostFtdcBillNameType [33]byte
type TThostFtdcPublishStatusType byte
type TThostFtdcEnumValueIDType [65]byte
type TThostFtdcEnumValueTypeType [33]byte
type TThostFtdcEnumValueLabelType [65]byte
type TThostFtdcEnumValueResultType [33]byte
type TThostFtdcSystemStatusType byte
type TThostFtdcSettlementStatusType byte
type TThostFtdcRangeIntTypeType [33]byte
type TThostFtdcRangeIntFromType [33]byte
type TThostFtdcRangeIntToType [33]byte
type TThostFtdcFunctionIDType [25]byte
type TThostFtdcFunctionValueCodeType [257]byte
type TThostFtdcFunctionNameType [65]byte
type TThostFtdcRoleIDType [11]byte
type TThostFtdcRoleNameType [41]byte
type TThostFtdcDescriptionType [401]byte
type TThostFtdcCombineIDType [25]byte
type TThostFtdcCombineTypeType [25]byte
type TThostFtdcInvestorTypeType byte
type TThostFtdcBrokerTypeType byte
type TThostFtdcRiskLevelType byte
type TThostFtdcFeeAcceptStyleType byte
type TThostFtdcPasswordTypeType byte
type TThostFtdcAlgorithmType byte
type TThostFtdcIncludeCloseProfitType byte
type TThostFtdcAllWithoutTradeType byte
type TThostFtdcCommentType [31]byte
type TThostFtdcVersionType [4]byte
type TThostFtdcTradeCodeType [7]byte
type TThostFtdcTradeDateType [9]byte
type TThostFtdcTradeTimeType [9]byte
type TThostFtdcTradeSerialType [9]byte
type TThostFtdcTradeSerialNoType int64
type TThostFtdcFutureIDType [11]byte
type TThostFtdcBankIDType [4]byte
type TThostFtdcBankBrchIDType [5]byte
type TThostFtdcBankBranchIDType [11]byte
type TThostFtdcOperNoType [17]byte
type TThostFtdcDeviceIDType [3]byte
type TThostFtdcRecordNumType [7]byte
type TThostFtdcFutureAccountType [22]byte
type TThostFtdcFuturePwdFlagType byte
type TThostFtdcTransferTypeType byte
type TThostFtdcFutureAccPwdType [17]byte
type TThostFtdcCurrencyCodeType [4]byte
type TThostFtdcRetCodeType [5]byte
type TThostFtdcRetInfoType [129]byte
type TThostFtdcTradeAmtType [20]byte
type TThostFtdcUseAmtType [20]byte
type TThostFtdcFetchAmtType [20]byte
type TThostFtdcTransferValidFlagType byte
type TThostFtdcCertCodeType [21]byte
type TThostFtdcReasonType byte
type TThostFtdcFundProjectIDType [5]byte
type TThostFtdcSexType byte
type TThostFtdcProfessionType [101]byte
type TThostFtdcNationalType [31]byte
type TThostFtdcProvinceType [51]byte
type TThostFtdcRegionType [16]byte
type TThostFtdcCountryType [16]byte
type TThostFtdcLicenseNOType [33]byte
type TThostFtdcCompanyTypeType [16]byte
type TThostFtdcBusinessScopeType [1001]byte
type TThostFtdcCapitalCurrencyType [4]byte
type TThostFtdcUserTypeType byte
type TThostFtdcBranchIDType [9]byte
type TThostFtdcRateTypeType byte
type TThostFtdcNoteTypeType byte
type TThostFtdcSettlementStyleType byte
type TThostFtdcBrokerDNSType [256]byte
type TThostFtdcSentenceType [501]byte
type TThostFtdcSettlementBillTypeType byte
type TThostFtdcUserRightTypeType byte
type TThostFtdcMarginPriceTypeType byte
type TThostFtdcBillGenStatusType byte
type TThostFtdcAlgoTypeType byte
type TThostFtdcHandlePositionAlgoIDType byte
type TThostFtdcFindMarginRateAlgoIDType byte
type TThostFtdcHandleTradingAccountAlgoIDType byte
type TThostFtdcPersonTypeType byte
type TThostFtdcQueryInvestorRangeType byte
type TThostFtdcInvestorRiskStatusType byte
type TThostFtdcLegIDType int64
type TThostFtdcLegMultipleType int64
type TThostFtdcImplyLevelType int64
type TThostFtdcClearAccountType [33]byte
type TThostFtdcOrganNOType [6]byte
type TThostFtdcClearbarchIDType [6]byte
type TThostFtdcUserEventTypeType byte
type TThostFtdcUserEventInfoType [1025]byte
type TThostFtdcCloseStyleType byte
type TThostFtdcStatModeType byte
type TThostFtdcParkedOrderStatusType byte
type TThostFtdcParkedOrderIDType [13]byte
type TThostFtdcParkedOrderActionIDType [13]byte
type TThostFtdcVirDealStatusType byte
type TThostFtdcOrgSystemIDType byte
type TThostFtdcVirTradeStatusType byte
type TThostFtdcVirBankAccTypeType byte
type TThostFtdcVirementStatusType byte
type TThostFtdcVirementAvailAbilityType byte
type TThostFtdcVirementTradeCodeType byte
type TThostFtdcPhotoTypeNameType [41]byte
type TThostFtdcPhotoTypeIDType [5]byte
type TThostFtdcPhotoNameType [161]byte
type TThostFtdcTopicIDType int64
type TThostFtdcReportTypeIDType [3]byte
type TThostFtdcCharacterIDType [5]byte
type TThostFtdcAMLParamIDType [21]byte
type TThostFtdcAMLInvestorTypeType [3]byte
type TThostFtdcAMLIdCardTypeType [3]byte
type TThostFtdcAMLTradeDirectType [3]byte
type TThostFtdcAMLTradeModelType [3]byte
type TThostFtdcAMLOpParamValueType float64
type TThostFtdcAMLCustomerCardTypeType [81]byte
type TThostFtdcAMLInstitutionNameType [65]byte
type TThostFtdcAMLDistrictIDType [7]byte
type TThostFtdcAMLRelationShipType [3]byte
type TThostFtdcAMLInstitutionTypeType [3]byte
type TThostFtdcAMLInstitutionIDType [13]byte
type TThostFtdcAMLAccountTypeType [5]byte
type TThostFtdcAMLTradingTypeType [7]byte
type TThostFtdcAMLTransactClassType [7]byte
type TThostFtdcAMLCapitalIOType [3]byte
type TThostFtdcAMLSiteType [10]byte
type TThostFtdcAMLCapitalPurposeType [129]byte
type TThostFtdcAMLReportTypeType [2]byte
type TThostFtdcAMLSerialNoType [5]byte
type TThostFtdcAMLStatusType [2]byte
type TThostFtdcAMLGenStatusType byte
type TThostFtdcAMLSeqCodeType [65]byte
type TThostFtdcAMLFileNameType [257]byte
type TThostFtdcAMLMoneyType float64
type TThostFtdcAMLFileAmountType int64
type TThostFtdcCFMMCKeyType [21]byte
type TThostFtdcCFMMCTokenType [21]byte
type TThostFtdcCFMMCKeyKindType byte
type TThostFtdcAMLReportNameType [81]byte
type TThostFtdcIndividualNameType [51]byte
type TThostFtdcCurrencyIDType [4]byte
type TThostFtdcCustNumberType [36]byte
type TThostFtdcOrganCodeType [36]byte
type TThostFtdcOrganNameType [71]byte
type TThostFtdcSuperOrganCodeType [12]byte
type TThostFtdcSubBranchIDType [31]byte
type TThostFtdcSubBranchNameType [71]byte
type TThostFtdcBranchNetCodeType [31]byte
type TThostFtdcBranchNetNameType [71]byte
type TThostFtdcOrganFlagType [2]byte
type TThostFtdcBankCodingForFutureType [33]byte
type TThostFtdcBankReturnCodeType [7]byte
type TThostFtdcPlateReturnCodeType [5]byte
type TThostFtdcBankSubBranchIDType [31]byte
type TThostFtdcFutureBranchIDType [31]byte
type TThostFtdcReturnCodeType [7]byte
type TThostFtdcOperatorCodeType [17]byte
type TThostFtdcClearDepIDType [6]byte
type TThostFtdcClearBrchIDType [6]byte
type TThostFtdcClearNameType [71]byte
type TThostFtdcBankAccountNameType [71]byte
type TThostFtdcInvDepIDType [6]byte
type TThostFtdcInvBrchIDType [6]byte
type TThostFtdcMessageFormatVersionType [36]byte
type TThostFtdcDigestType [36]byte
type TThostFtdcAuthenticDataType [129]byte
type TThostFtdcPasswordKeyType [129]byte
type TThostFtdcFutureAccountNameType [129]byte
type TThostFtdcMobilePhoneType [21]byte
type TThostFtdcFutureMainKeyType [129]byte
type TThostFtdcFutureWorkKeyType [129]byte
type TThostFtdcFutureTransKeyType [129]byte
type TThostFtdcBankMainKeyType [129]byte
type TThostFtdcBankWorkKeyType [129]byte
type TThostFtdcBankTransKeyType [129]byte
type TThostFtdcBankServerDescriptionType [129]byte
type TThostFtdcAddInfoType [129]byte
type TThostFtdcDescrInfoForReturnCodeType [129]byte
type TThostFtdcCountryCodeType [21]byte
type TThostFtdcSerialType int64
type TThostFtdcPlateSerialType int64
type TThostFtdcBankSerialType [13]byte
type TThostFtdcCorrectSerialType int64
type TThostFtdcFutureSerialType int64
type TThostFtdcApplicationIDType int64
type TThostFtdcBankProxyIDType int64
type TThostFtdcFBTCoreIDType int64
type TThostFtdcServerPortType int64
type TThostFtdcRepealedTimesType int64
type TThostFtdcRepealTimeIntervalType int64
type TThostFtdcTotalTimesType int64
type TThostFtdcFBTRequestIDType int64
type TThostFtdcTIDType int64
type TThostFtdcTradeAmountType float64
type TThostFtdcCustFeeType float64
type TThostFtdcFutureFeeType float64
type TThostFtdcSingleMaxAmtType float64
type TThostFtdcSingleMinAmtType float64
type TThostFtdcTotalAmtType float64
type TThostFtdcCertificationTypeType byte
type TThostFtdcFileBusinessCodeType byte
type TThostFtdcCashExchangeCodeType byte
type TThostFtdcYesNoIndicatorType byte
type TThostFtdcBanlanceTypeType byte
type TThostFtdcGenderType byte
type TThostFtdcFeePayFlagType byte
type TThostFtdcPassWordKeyTypeType byte
type TThostFtdcFBTPassWordTypeType byte
type TThostFtdcFBTEncryModeType byte
type TThostFtdcBankRepealFlagType byte
type TThostFtdcBrokerRepealFlagType byte
type TThostFtdcInstitutionTypeType byte
type TThostFtdcLastFragmentType byte
type TThostFtdcBankAccStatusType byte
type TThostFtdcMoneyAccountStatusType byte
type TThostFtdcManageStatusType byte
type TThostFtdcSystemTypeType byte
type TThostFtdcTxnEndFlagType byte
type TThostFtdcProcessStatusType byte
type TThostFtdcCustTypeType byte
type TThostFtdcFBTTransferDirectionType byte
type TThostFtdcOpenOrDestroyType byte
type TThostFtdcAvailabilityFlagType byte
type TThostFtdcOrganTypeType byte
type TThostFtdcOrganLevelType byte
type TThostFtdcProtocalIDType byte
type TThostFtdcConnectModeType byte
type TThostFtdcSyncModeType byte
type TThostFtdcBankAccTypeType byte
type TThostFtdcFutureAccTypeType byte
type TThostFtdcOrganStatusType byte
type TThostFtdcCCBFeeModeType byte
type TThostFtdcCommApiTypeType byte
type TThostFtdcServiceIDType int64
type TThostFtdcServiceLineNoType int64
type TThostFtdcServiceNameType [61]byte
type TThostFtdcLinkStatusType byte
type TThostFtdcCommApiPointerType int64
type TThostFtdcPwdFlagType byte
type TThostFtdcSecuAccTypeType byte
type TThostFtdcTransferStatusType byte
type TThostFtdcSponsorTypeType byte
type TThostFtdcReqRspTypeType byte
type TThostFtdcFBTUserEventTypeType byte
type TThostFtdcBankIDByBankType [21]byte
type TThostFtdcBankOperNoType [4]byte
type TThostFtdcBankCustNoType [21]byte
type TThostFtdcDBOPSeqNoType int64
type TThostFtdcTableNameType [61]byte
type TThostFtdcPKNameType [201]byte
type TThostFtdcPKValueType [501]byte
type TThostFtdcDBOperationType byte
type TThostFtdcSyncFlagType byte
type TThostFtdcTargetIDType [4]byte
type TThostFtdcSyncTypeType byte
type TThostFtdcFBETimeType [7]byte
type TThostFtdcFBEBankNoType [13]byte
type TThostFtdcFBECertNoType [13]byte
type TThostFtdcExDirectionType byte
type TThostFtdcFBEBankAccountType [33]byte
type TThostFtdcFBEBankAccountNameType [61]byte
type TThostFtdcFBEAmtType float64
type TThostFtdcFBEBusinessTypeType [3]byte
type TThostFtdcFBEPostScriptType [61]byte
type TThostFtdcFBERemarkType [71]byte
type TThostFtdcExRateType float64
type TThostFtdcFBEResultFlagType byte
type TThostFtdcFBERtnMsgType [61]byte
type TThostFtdcFBEExtendMsgType [61]byte
type TThostFtdcFBEBusinessSerialType [31]byte
type TThostFtdcFBESystemSerialType [21]byte
type TThostFtdcFBETotalExCntType int64
type TThostFtdcFBEExchStatusType byte
type TThostFtdcFBEFileFlagType byte
type TThostFtdcFBEAlreadyTradeType byte
type TThostFtdcFBEOpenBankType [61]byte
type TThostFtdcFBEUserEventTypeType byte
type TThostFtdcFBEFileNameType [21]byte
type TThostFtdcFBEBatchSerialType [21]byte
type TThostFtdcFBEReqFlagType byte
type TThostFtdcNotifyClassType byte
type TThostFtdcRiskNofityInfoType [257]byte
type TThostFtdcForceCloseSceneIdType [24]byte
type TThostFtdcForceCloseTypeType byte
type TThostFtdcInstrumentIDsType [101]byte
type TThostFtdcRiskNotifyMethodType byte
type TThostFtdcRiskNotifyStatusType byte
type TThostFtdcRiskUserEventType byte
type TThostFtdcParamIDType int64
type TThostFtdcParamNameType [41]byte
type TThostFtdcParamValueType [41]byte
type TThostFtdcConditionalOrderSortTypeType byte
type TThostFtdcSendTypeType byte
type TThostFtdcClientIDStatusType byte
type TThostFtdcIndustryIDType [17]byte
type TThostFtdcQuestionIDType [5]byte
type TThostFtdcQuestionContentType [41]byte
type TThostFtdcOptionIDType [13]byte
type TThostFtdcOptionContentType [61]byte
type TThostFtdcQuestionTypeType byte
type TThostFtdcProcessIDType [33]byte
type TThostFtdcSeqNoType int64
type TThostFtdcUOAProcessStatusType [3]byte
type TThostFtdcProcessTypeType [3]byte
type TThostFtdcBusinessTypeType byte
type TThostFtdcCfmmcReturnCodeType byte
type TThostFtdcExReturnCodeType int64
type TThostFtdcClientTypeType byte
type TThostFtdcExchangeIDTypeType byte
type TThostFtdcExClientIDTypeType byte
type TThostFtdcClientClassifyType [11]byte
type TThostFtdcUOAOrganTypeType [11]byte
type TThostFtdcUOACountryCodeType [11]byte
type TThostFtdcAreaCodeType [11]byte
type TThostFtdcFuturesIDType [21]byte
type TThostFtdcCffmcDateType [11]byte
type TThostFtdcCffmcTimeType [11]byte
type TThostFtdcNocIDType [21]byte
type TThostFtdcUpdateFlagType byte
type TThostFtdcApplyOperateIDType byte
type TThostFtdcApplyStatusIDType byte
type TThostFtdcSendMethodType byte
type TThostFtdcEventTypeType [33]byte
type TThostFtdcEventModeType byte
type TThostFtdcUOAAutoSendType byte
type TThostFtdcQueryDepthType int64
type TThostFtdcDataCenterIDType int64
type TThostFtdcFlowIDType byte
type TThostFtdcCheckLevelType byte
type TThostFtdcCheckNoType int64
type TThostFtdcCheckStatusType byte
type TThostFtdcUsedStatusType byte
type TThostFtdcRateTemplateNameType [61]byte
type TThostFtdcPropertyStringType [2049]byte
type TThostFtdcBankAcountOriginType byte
type TThostFtdcMonthBillTradeSumType byte
type TThostFtdcFBTTradeCodeEnumType byte
type TThostFtdcRateTemplateIDType [9]byte
type TThostFtdcRiskRateType [21]byte
type TThostFtdcTimestampType int64
type TThostFtdcInvestorIDRuleNameType [61]byte
type TThostFtdcInvestorIDRuleExprType [513]byte
type TThostFtdcLastDriftType int64
type TThostFtdcLastSuccessType int64
type TThostFtdcAuthKeyType [41]byte
type TThostFtdcSerialNumberType [17]byte
type TThostFtdcOTPTypeType byte
type TThostFtdcOTPVendorsIDType [2]byte
type TThostFtdcOTPVendorsNameType [61]byte
type TThostFtdcOTPStatusType byte
type TThostFtdcBrokerUserTypeType byte
type TThostFtdcFutureTypeType byte
type TThostFtdcFundEventTypeType byte
type TThostFtdcAccountSourceTypeType byte
type TThostFtdcCodeSourceTypeType byte
type TThostFtdcUserRangeType byte
type TThostFtdcTimeSpanType [9]byte
type TThostFtdcImportSequenceIDType [17]byte
type TThostFtdcByGroupType byte
type TThostFtdcTradeSumStatModeType byte
type TThostFtdcComTypeType int64
type TThostFtdcUserProductIDType [33]byte
type TThostFtdcUserProductNameType [65]byte
type TThostFtdcUserProductMemoType [129]byte
type TThostFtdcCSRCCancelFlagType [2]byte
type TThostFtdcCSRCDateType [11]byte
type TThostFtdcCSRCInvestorNameType [201]byte
type TThostFtdcCSRCOpenInvestorNameType [101]byte
type TThostFtdcCSRCInvestorIDType [13]byte
type TThostFtdcCSRCIdentifiedCardNoType [51]byte
type TThostFtdcCSRCClientIDType [11]byte
type TThostFtdcCSRCBankFlagType [3]byte
type TThostFtdcCSRCBankAccountType [23]byte
type TThostFtdcCSRCOpenNameType [401]byte
type TThostFtdcCSRCMemoType [101]byte
type TThostFtdcCSRCTimeType [11]byte
type TThostFtdcCSRCTradeIDType [21]byte
type TThostFtdcCSRCExchangeInstIDType [31]byte
type TThostFtdcCSRCMortgageNameType [7]byte
type TThostFtdcCSRCReasonType [3]byte
type TThostFtdcIsSettlementType [2]byte
type TThostFtdcCSRCMoneyType float64
type TThostFtdcCSRCPriceType float64
type TThostFtdcCSRCOptionsTypeType [2]byte
type TThostFtdcCSRCStrikePriceType float64
type TThostFtdcCSRCTargetProductIDType [3]byte
type TThostFtdcCSRCTargetInstrIDType [31]byte
type TThostFtdcCommModelNameType [161]byte
type TThostFtdcCommModelMemoType [1025]byte
type TThostFtdcExprSetModeType byte
type TThostFtdcRateInvestorRangeType byte
type TThostFtdcAgentBrokerIDType [13]byte
type TThostFtdcDRIdentityIDType int64
type TThostFtdcDRIdentityNameType [65]byte
type TThostFtdcDBLinkIDType [31]byte
type TThostFtdcSyncDataStatusType byte
type TThostFtdcTradeSourceType byte
type TThostFtdcFlexStatModeType byte
type TThostFtdcByInvestorRangeType byte
type TThostFtdcSRiskRateType [21]byte
type TThostFtdcSequenceNo12Type int64
type TThostFtdcPropertyInvestorRangeType byte
type TThostFtdcFileStatusType byte
type TThostFtdcFileGenStyleType byte
type TThostFtdcSysOperModeType byte
type TThostFtdcSysOperTypeType byte
type TThostFtdcCSRCDataQueyTypeType byte
type TThostFtdcFreezeStatusType byte
type TThostFtdcStandardStatusType byte
type TThostFtdcCSRCFreezeStatusType [2]byte
type TThostFtdcRightParamTypeType byte
type TThostFtdcRightTemplateIDType [9]byte
type TThostFtdcRightTemplateNameType [61]byte
type TThostFtdcDataStatusType byte
type TThostFtdcAMLCheckStatusType byte
type TThostFtdcAmlDateTypeType byte
type TThostFtdcAmlCheckLevelType byte
type TThostFtdcAmlCheckFlowType [2]byte
type TThostFtdcDataTypeType [129]byte
type TThostFtdcExportFileTypeType byte
type TThostFtdcSettleManagerTypeType byte
type TThostFtdcSettleManagerIDType [33]byte
type TThostFtdcSettleManagerNameType [129]byte
type TThostFtdcSettleManagerLevelType byte
type TThostFtdcSettleManagerGroupType byte
type TThostFtdcCheckResultMemoType [1025]byte
type TThostFtdcFunctionUrlType [1025]byte
type TThostFtdcAuthInfoType [129]byte
type TThostFtdcAuthCodeType [17]byte
type TThostFtdcLimitUseTypeType byte
type TThostFtdcDataResourceType byte
type TThostFtdcMarginTypeType byte
type TThostFtdcActiveTypeType byte
type TThostFtdcMarginRateTypeType byte
type TThostFtdcBackUpStatusType byte
type TThostFtdcInitSettlementType byte
type TThostFtdcReportStatusType byte
type TThostFtdcSaveStatusType byte
type TThostFtdcSettArchiveStatusType byte
type TThostFtdcCTPTypeType byte
type TThostFtdcToolIDType [9]byte
type TThostFtdcToolNameType [81]byte
type TThostFtdcCloseDealTypeType byte
type TThostFtdcMortgageFundUseRangeType byte
type TThostFtdcCurrencyUnitType float64
type TThostFtdcExchangeRateType float64
type TThostFtdcSpecProductTypeType byte
type TThostFtdcFundMortgageTypeType byte
type TThostFtdcAccountSettlementParamIDType byte
type TThostFtdcCurrencyNameType [31]byte
type TThostFtdcCurrencySignType [4]byte
type TThostFtdcFundMortDirectionType byte
type TThostFtdcBusinessClassType byte
type TThostFtdcSwapSourceTypeType byte
type TThostFtdcCurrExDirectionType byte
type TThostFtdcCurrencySwapStatusType byte
type TThostFtdcCurrExchCertNoType [13]byte
type TThostFtdcBatchSerialNoType [21]byte
type TThostFtdcReqFlagType byte
type TThostFtdcResFlagType byte
type TThostFtdcPageControlType [2]byte
type TThostFtdcRecordCountType int64
type TThostFtdcCurrencySwapMemoType [101]byte
type TThostFtdcExStatusType byte
type TThostFtdcClientRegionType byte
type TThostFtdcWorkPlaceType [101]byte
type TThostFtdcBusinessPeriodType [21]byte
type TThostFtdcWebSiteType [101]byte
type TThostFtdcUOAIdCardTypeType [3]byte
type TThostFtdcClientModeType [3]byte
type TThostFtdcInvestorFullNameType [101]byte
type TThostFtdcUOABrokerIDType [11]byte
type TThostFtdcUOAZipCodeType [11]byte
type TThostFtdcUOAEMailType [101]byte
type TThostFtdcOldCityType [41]byte
type TThostFtdcCorporateIdentifiedCardNoType [101]byte
type TThostFtdcHasBoardType byte
type TThostFtdcStartModeType byte
type TThostFtdcTemplateTypeType byte
type TThostFtdcLoginModeType byte
type TThostFtdcPromptTypeType byte
type TThostFtdcLedgerManageIDType [51]byte
type TThostFtdcInvestVarietyType [101]byte
type TThostFtdcBankAccountTypeType [2]byte
type TThostFtdcLedgerManageBankType [101]byte
type TThostFtdcCffexDepartmentNameType [101]byte
type TThostFtdcCffexDepartmentCodeType [9]byte
type TThostFtdcHasTrusteeType byte
type TThostFtdcCSRCMemo1Type [41]byte
type TThostFtdcAssetmgrCFullNameType [101]byte
type TThostFtdcAssetmgrApprovalNOType [51]byte
type TThostFtdcAssetmgrMgrNameType [401]byte
type TThostFtdcAmTypeType byte
type TThostFtdcCSRCAmTypeType [5]byte
type TThostFtdcCSRCFundIOTypeType byte
type TThostFtdcCusAccountTypeType byte
type TThostFtdcCSRCNationalType [4]byte
type TThostFtdcCSRCSecAgentIDType [11]byte
type TThostFtdcLanguageTypeType byte
type TThostFtdcAmAccountType [23]byte
type TThostFtdcAssetmgrClientTypeType byte
type TThostFtdcAssetmgrTypeType byte
type TThostFtdcUOMType [11]byte
type TThostFtdcSHFEInstLifePhaseType [3]byte
type TThostFtdcSHFEProductClassType [11]byte
type TThostFtdcPriceDecimalType [2]byte
type TThostFtdcInTheMoneyFlagType [2]byte
type TThostFtdcCheckInstrTypeType byte
type TThostFtdcDeliveryTypeType byte
type TThostFtdcBigMoneyType float64
type TThostFtdcMaxMarginSideAlgorithmType byte
type TThostFtdcDAClientTypeType byte
type TThostFtdcCombinInstrIDType [61]byte
type TThostFtdcCombinSettlePriceType [61]byte
type TThostFtdcDCEPriorityType int64
type TThostFtdcTradeGroupIDType int64
type TThostFtdcIsCheckPrepaType int64
type TThostFtdcUOAAssetmgrTypeType byte
type TThostFtdcDirectionEnType byte
type TThostFtdcOffsetFlagEnType byte
type TThostFtdcHedgeFlagEnType byte
type TThostFtdcFundIOTypeEnType byte
type TThostFtdcFundTypeEnType byte
type TThostFtdcFundDirectionEnType byte
type TThostFtdcFundMortDirectionEnType byte
type TThostFtdcSwapBusinessTypeType [3]byte
type TThostFtdcOptionsTypeType byte
type TThostFtdcStrikeModeType byte
type TThostFtdcStrikeTypeType byte
type TThostFtdcApplyTypeType byte
type TThostFtdcGiveUpDataSourceType byte
type TThostFtdcExecOrderSysIDType [21]byte
type TThostFtdcExecResultType byte
type TThostFtdcStrikeSequenceType int64
type TThostFtdcStrikeTimeType [13]byte
type TThostFtdcCombinationTypeType byte
type TThostFtdcOptionRoyaltyPriceTypeType byte
type TThostFtdcBalanceAlgorithmType byte
type TThostFtdcActionTypeType byte
type TThostFtdcForQuoteStatusType byte
type TThostFtdcValueMethodType byte
type TThostFtdcExecOrderPositionFlagType byte
type TThostFtdcExecOrderCloseFlagType byte
type TThostFtdcProductTypeType byte
type TThostFtdcCZCEUploadFileNameType byte
type TThostFtdcDCEUploadFileNameType byte
type TThostFtdcSHFEUploadFileNameType byte
type TThostFtdcCFFEXUploadFileNameType byte
type TThostFtdcCombDirectionType byte
type TThostFtdcStrikeOffsetTypeType byte
type TThostFtdcReserveOpenAccStasType byte
type TThostFtdcLoginRemarkType [36]byte
type TThostFtdcInvestUnitIDType [17]byte
type TThostFtdcBulletinIDType int64
type TThostFtdcNewsTypeType [3]byte
type TThostFtdcNewsUrgencyType byte
type TThostFtdcAbstractType [81]byte
type TThostFtdcComeFromType [21]byte
type TThostFtdcURLLinkType [201]byte
type TThostFtdcLongIndividualNameType [161]byte
type TThostFtdcLongFBEBankAccountNameType [161]byte
type TThostFtdcDateTimeType [17]byte
type TThostFtdcWeakPasswordSourceType byte
type TThostFtdcRandomStringType [17]byte
type TThostFtdcOptSelfCloseFlagType byte
type TThostFtdcBizTypeType byte
