import api from './serviceApi';

const VALIDATE_PATH = '/validator/v1/validate';
const SAVE_PATH = '/validator/v1/save';
const GETALL_PATH = '/validator/v1/getDocuments';
const MOVE_TO_BLACKLIST_PATH = '/validator/v1/blacklist';
const DELETE_PATH = '/validator/v1/delete';
const GETBLACKLIST_PATH = '/validator/v1/getBlacklist';

export const validateDocument = async (documentNumber) => {
  const header = { headers: { 'Content-Type': 'application/json' } };
  const result = await api.post(`${VALIDATE_PATH}`, {number: documentNumber}, header);
  return result;
}

export const saveDocument = async (documentNumber, documentType) => {
  const header = { headers: { 'Content-Type': 'application/json' } };
  const request = JSON.stringify({
    number: `${documentNumber}`,
    type: `${documentType}`,
    blacklist: false,
  });
  const result = await api.post(`${SAVE_PATH}`, request, header);
  return result;
}

export const getAllDocuments = async () => {
  const header = { headers: { 'Content-Type': 'application/json' } };
  const result = await api.get(`${GETALL_PATH}`, header);
  return result;
}

export const moveToBlacklist = async (documentId) => {
  const header = { headers: { 'Content-Type': 'application/json' } };
  const result = await api.get(`${MOVE_TO_BLACKLIST_PATH}/${documentId}`, header);
  return result;
}

export const deleteDocument = async (documentId) => {
  const header = { headers: { 'Content-Type': 'application/json' } };
  const result = await api.delete(`${DELETE_PATH}/${documentId}`, header);
  return result;
}


export const getBlacklist = async () => {
  const header = { headers: { 'Content-Type': 'application/json' } };
  const result = await api.get(`${GETBLACKLIST_PATH}`, header);
  return result;
}

export default {
  validateDocument,
  saveDocument,
  getAllDocuments,
  moveToBlacklist
  };
  