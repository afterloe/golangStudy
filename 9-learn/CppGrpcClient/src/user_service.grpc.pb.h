// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: user_service.proto
#ifndef GRPC_user_5fservice_2eproto__INCLUDED
#define GRPC_user_5fservice_2eproto__INCLUDED

#include "user_service.pb.h"

#include <functional>
#include <grpcpp/generic/async_generic_service.h>
#include <grpcpp/support/async_stream.h>
#include <grpcpp/support/async_unary_call.h>
#include <grpcpp/support/client_callback.h>
#include <grpcpp/client_context.h>
#include <grpcpp/completion_queue.h>
#include <grpcpp/support/message_allocator.h>
#include <grpcpp/support/method_handler.h>
#include <grpcpp/impl/proto_utils.h>
#include <grpcpp/impl/rpc_method.h>
#include <grpcpp/support/server_callback.h>
#include <grpcpp/impl/server_callback_handlers.h>
#include <grpcpp/server_context.h>
#include <grpcpp/impl/service_type.h>
#include <grpcpp/support/status.h>
#include <grpcpp/support/stub_options.h>
#include <grpcpp/support/sync_stream.h>

namespace Logic {

class UserServer final {
 public:
  static constexpr char const* service_full_name() {
    return "Logic.UserServer";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    virtual ::grpc::Status Login(::grpc::ClientContext* context, const ::Logic::LoginRequest& request, ::Logic::LoginResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::Logic::LoginResponse>> AsyncLogin(::grpc::ClientContext* context, const ::Logic::LoginRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::Logic::LoginResponse>>(AsyncLoginRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::Logic::LoginResponse>> PrepareAsyncLogin(::grpc::ClientContext* context, const ::Logic::LoginRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::Logic::LoginResponse>>(PrepareAsyncLoginRaw(context, request, cq));
    }
    class async_interface {
     public:
      virtual ~async_interface() {}
      virtual void Login(::grpc::ClientContext* context, const ::Logic::LoginRequest* request, ::Logic::LoginResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void Login(::grpc::ClientContext* context, const ::Logic::LoginRequest* request, ::Logic::LoginResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
    };
    typedef class async_interface experimental_async_interface;
    virtual class async_interface* async() { return nullptr; }
    class async_interface* experimental_async() { return async(); }
   private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::Logic::LoginResponse>* AsyncLoginRaw(::grpc::ClientContext* context, const ::Logic::LoginRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::Logic::LoginResponse>* PrepareAsyncLoginRaw(::grpc::ClientContext* context, const ::Logic::LoginRequest& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());
    ::grpc::Status Login(::grpc::ClientContext* context, const ::Logic::LoginRequest& request, ::Logic::LoginResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::Logic::LoginResponse>> AsyncLogin(::grpc::ClientContext* context, const ::Logic::LoginRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::Logic::LoginResponse>>(AsyncLoginRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::Logic::LoginResponse>> PrepareAsyncLogin(::grpc::ClientContext* context, const ::Logic::LoginRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::Logic::LoginResponse>>(PrepareAsyncLoginRaw(context, request, cq));
    }
    class async final :
      public StubInterface::async_interface {
     public:
      void Login(::grpc::ClientContext* context, const ::Logic::LoginRequest* request, ::Logic::LoginResponse* response, std::function<void(::grpc::Status)>) override;
      void Login(::grpc::ClientContext* context, const ::Logic::LoginRequest* request, ::Logic::LoginResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
     private:
      friend class Stub;
      explicit async(Stub* stub): stub_(stub) { }
      Stub* stub() { return stub_; }
      Stub* stub_;
    };
    class async* async() override { return &async_stub_; }

   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    class async async_stub_{this};
    ::grpc::ClientAsyncResponseReader< ::Logic::LoginResponse>* AsyncLoginRaw(::grpc::ClientContext* context, const ::Logic::LoginRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::Logic::LoginResponse>* PrepareAsyncLoginRaw(::grpc::ClientContext* context, const ::Logic::LoginRequest& request, ::grpc::CompletionQueue* cq) override;
    const ::grpc::internal::RpcMethod rpcmethod_Login_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    virtual ::grpc::Status Login(::grpc::ServerContext* context, const ::Logic::LoginRequest* request, ::Logic::LoginResponse* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_Login : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_Login() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_Login() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Login(::grpc::ServerContext* /*context*/, const ::Logic::LoginRequest* /*request*/, ::Logic::LoginResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestLogin(::grpc::ServerContext* context, ::Logic::LoginRequest* request, ::grpc::ServerAsyncResponseWriter< ::Logic::LoginResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_Login<Service > AsyncService;
  template <class BaseClass>
  class WithCallbackMethod_Login : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_Login() {
      ::grpc::Service::MarkMethodCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::Logic::LoginRequest, ::Logic::LoginResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::Logic::LoginRequest* request, ::Logic::LoginResponse* response) { return this->Login(context, request, response); }));}
    void SetMessageAllocatorFor_Login(
        ::grpc::MessageAllocator< ::Logic::LoginRequest, ::Logic::LoginResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(0);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::Logic::LoginRequest, ::Logic::LoginResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_Login() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Login(::grpc::ServerContext* /*context*/, const ::Logic::LoginRequest* /*request*/, ::Logic::LoginResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* Login(
      ::grpc::CallbackServerContext* /*context*/, const ::Logic::LoginRequest* /*request*/, ::Logic::LoginResponse* /*response*/)  { return nullptr; }
  };
  typedef WithCallbackMethod_Login<Service > CallbackService;
  typedef CallbackService ExperimentalCallbackService;
  template <class BaseClass>
  class WithGenericMethod_Login : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_Login() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_Login() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Login(::grpc::ServerContext* /*context*/, const ::Logic::LoginRequest* /*request*/, ::Logic::LoginResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithRawMethod_Login : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_Login() {
      ::grpc::Service::MarkMethodRaw(0);
    }
    ~WithRawMethod_Login() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Login(::grpc::ServerContext* /*context*/, const ::Logic::LoginRequest* /*request*/, ::Logic::LoginResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestLogin(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_Login : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_Login() {
      ::grpc::Service::MarkMethodRawCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->Login(context, request, response); }));
    }
    ~WithRawCallbackMethod_Login() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Login(::grpc::ServerContext* /*context*/, const ::Logic::LoginRequest* /*request*/, ::Logic::LoginResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* Login(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_Login : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_Login() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::internal::StreamedUnaryHandler<
          ::Logic::LoginRequest, ::Logic::LoginResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::Logic::LoginRequest, ::Logic::LoginResponse>* streamer) {
                       return this->StreamedLogin(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_Login() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status Login(::grpc::ServerContext* /*context*/, const ::Logic::LoginRequest* /*request*/, ::Logic::LoginResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedLogin(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::Logic::LoginRequest,::Logic::LoginResponse>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_Login<Service > StreamedUnaryService;
  typedef Service SplitStreamedService;
  typedef WithStreamedUnaryMethod_Login<Service > StreamedService;
};

}  // namespace Logic


#endif  // GRPC_user_5fservice_2eproto__INCLUDED
